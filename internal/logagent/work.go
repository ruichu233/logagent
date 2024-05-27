package logagent

import (
	"context"
	"github.com/hpcloud/tail"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Work struct {
	ctx    context.Context
	cancel context.CancelFunc
	writer *kafka.Writer
	reader *tail.Tail
	buffer chan *kafka.Message
}

func NewWork(ctx context.Context, writer *kafka.Writer, tail *tail.Tail, bufferSize int) *Work {
	cctx, cancel := context.WithCancel(ctx)
	return &Work{
		ctx:    cctx,
		cancel: cancel,
		writer: writer,
		reader: tail,
		buffer: make(chan *kafka.Message, bufferSize),
	}
}

func (w *Work) run() {
	go w.ReadFromTail()
	go w.WriteToKafka()
}

func (w *Work) close() {
	time.Sleep(5 * time.Second)
	w.cancel()
	w.reader.Cleanup()
	_ = w.writer.Close()
}

func (w *Work) ReadFromTail() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case line, ok := <-w.reader.Lines:
			if !ok {
				break
			}
			logrus.Infoln("read from tail succ,val= ", line.Text)
			w.buffer <- &kafka.Message{
				Key:   []byte(strconv.FormatInt(line.Time.Unix(), 10)),
				Value: []byte(line.Text),
			}
		}
	}
}

func (w *Work) WriteToKafka() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case msg, ok := <-w.buffer:
			if !ok && msg == nil {
				break
			}
			if err := w.writer.WriteMessages(w.ctx, *msg); err != nil {
				logrus.Warnf("failed to write messages: %v", err)
				continue
			}
			logrus.Infoln("write to kafka succ,val= ", msg.Value)
		}
	}
}
