package logagent

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"os/signal"
	"syscall"
)

func NewLogAgentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "logagent",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	return cmd
}

func run() error {
	options, err := GetOptions()
	if err != nil {
		return err
	}
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   options.EtcdOptions.Endpoints,
		DialTimeout: options.EtcdOptions.DialTimeout,
	})
	if err != nil {
		return err
	}
	server, err := New(options, etcd)
	if err != nil {
		return err
	}
	if err := server.Run(); err != nil {
		return err
	}

	// 等待中断信号优雅关闭服务器
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 开始阻塞
	<-quit
	_ = server.Stop()

	logrus.Info("Shutdown Server ...")
	return nil
}
