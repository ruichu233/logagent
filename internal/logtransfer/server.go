package logtransfer

import (
	"context"
	"github.com/ruichu233/logagent/internal/pkg/read"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ctx     context.Context
	cancel  context.CancelFunc
	stopped bool
	reader  *read.Reader
	es      *ESClient
	ch      chan interface{}
}

func NewServer(opts *Options) *Server {
	reader := read.NewReader(opts.KafKaOptions)
	es := NewESClient(opts.ESOptions, opts.Index)
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		ctx:    ctx,
		cancel: cancel,
		reader: reader,
		es:     es,
		ch:     make(chan interface{}, opts.BufferSize),
	}
}

func (s *Server) Run() {
	go s.reader.Read(s.ctx, s.ch)
	go s.es.sentToES(s.ctx, s.ch)
}

func (s *Server) Stop() {
	if s.stopped {
		return
	}
	logrus.Info("server closing")
	s.cancel()
	s.reader.Close()
	s.stopped = true
	logrus.Info("server closed")
}
