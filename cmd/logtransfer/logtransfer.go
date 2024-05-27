package main

import (
	"github.com/ruichu233/logagent/internal/logtransfer"
	"os"
)

func main() {
	cmd := logtransfer.NewLogTransferCommend()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
