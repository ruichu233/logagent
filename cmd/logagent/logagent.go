package main

import (
	"github.com/ruichu233/logagent/internal/logagent"
	"os"
)

func main() {
	cmd := logagent.NewLogAgentCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
