package logtransfer

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func NewLogTransferCommend() *cobra.Command {
	cmd := &cobra.Command{
		Use: "logTransfor",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	return cmd
}

func run() error {

	opts, err := GetOptions()
	if err != nil {
		return err
	}
	server := NewServer(opts)
	server.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.Stop()
	return nil
}
