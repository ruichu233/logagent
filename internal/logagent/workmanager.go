package logagent

import (
	"context"
	"github.com/hpcloud/tail"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type WorkManager struct {
	ctx        context.Context
	opts       *Options
	confCentre *ConfCentre
	workerMap  map[string]*Work
	mu         sync.RWMutex
}

func NewWorkManager(ctx context.Context, etcd *clientv3.Client, opts *Options) (*WorkManager, error) {

	confCentre, err := NewConfCentre(make(chan struct{}), etcd)
	if err != nil {
		return nil, err
	}
	w := &WorkManager{
		ctx:        ctx,
		opts:       opts,
		confCentre: confCentre,
		workerMap:  make(map[string]*Work, len(confCentre.collects)),
	}
	for _, c := range w.confCentre.collects {
		work, err := w.getWorkerFromConf(c.Path, c.Topic)
		if err != nil {
			logrus.Warnf("get worker from config failed: %v\n", err)
		}
		w.AddWork(work)
	}

	go w.watch(ctx)
	return w, nil
}

func (w *WorkManager) AddWork(work *Work) {

	w.mu.Lock()
	defer w.mu.Unlock()
	w.workerMap[work.reader.Filename] = work
}

func (w *WorkManager) GetWork(filename string) (*Work, bool) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	work, ok := w.workerMap[filename]
	return work, ok
}

func (w *WorkManager) run() {
	for _, work := range w.workerMap {
		go work.run()
	}
}

func (w *WorkManager) stop() {
	for _, work := range w.workerMap {
		work.close()
	}
}

func (w *WorkManager) watch(ctx context.Context) {

	select {
	case <-ctx.Done():
		return
	// 配置发生改变
	case <-w.confCentre.updateChan:
		newConf := w.confCentre.collects
		for _, c := range newConf {
			// 未变
			if v, ok := w.workerMap[c.Path]; ok && v.writer.Topic == c.Topic {
				continue
			}
			// 新增
			newWork, err := w.getWorkerFromConf(c.Path, c.Topic)
			if err != nil {
				logrus.Warnf("get worker from config failed: %v\n", err)
				continue
			}
			w.AddWork(newWork)
			newWork.run()
		}
		// 处理删除部分
		for path, work := range w.workerMap {
			deleted := true
			for _, c := range newConf {
				if path == c.Path {
					deleted = false
					break
				}
			}

			if deleted {
				work.close()
				delete(w.workerMap, path)
			}
		}
	}

}

func (w *WorkManager) getWorkerFromConf(path, topic string) (*Work, error) {
	tailFile, err := tail.TailFile(path, tail.Config{
		Location: &tail.SeekInfo{
			Offset: w.opts.TailFileOptions.Location.Offset,
			Whence: w.opts.TailFileOptions.Location.Whence,
		},
		ReOpen:    w.opts.TailFileOptions.ReOpen,
		MustExist: w.opts.TailFileOptions.MustExist,
		Poll:      w.opts.TailFileOptions.Poll,
		Follow:    w.opts.TailFileOptions.Follow,
	})
	if err != nil {
		return nil, err
	}
	writer := &kafka.Writer{
		Addr:         kafka.TCP(w.opts.KafkaOptions.Brokers...),
		Topic:        topic,
		Balancer:     &kafka.Hash{},
		BatchSize:    w.opts.KafkaOptions.WriterOptions.BatchSize,
		BatchTimeout: w.opts.KafkaOptions.WriterOptions.BatchTimeout,
		MaxAttempts:  w.opts.KafkaOptions.WriterOptions.MaxAttempts,
		RequiredAcks: kafka.RequiredAcks(w.opts.KafkaOptions.WriterOptions.RequiredAcks),
		Async:        w.opts.KafkaOptions.WriterOptions.Async,
	}
	nw := NewWork(w.ctx, writer, tailFile, w.opts.BuffSize)
	return nw, nil
}
