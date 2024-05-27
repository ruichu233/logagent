package logagent

import (
	"context"
	"errors"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	ctx         context.Context
	stop        context.CancelFunc
	stopped     bool
	workManager *WorkManager
}

func New(opts *Options, etcd *clientv3.Client) (*Server, error) {
	ctx, cancel := context.WithCancel(context.Background())
	workManager, err := NewWorkManager(ctx, etcd, opts)
	if err != nil {
		cancel()
		return nil, err
	}
	server := &Server{
		ctx:         ctx,
		workManager: workManager,
		stop:        cancel,
	}
	return server, nil
}

func (s *Server) Run() error {
	if s.stopped {
		return errors.New("server is stopped")
	}
	s.workManager.run()
	return nil
}

func (s *Server) Stop() error {
	if s.stopped {
		return errors.New("server is stopped")
	}
	s.stop()
	s.stopped = true
	return nil
}
