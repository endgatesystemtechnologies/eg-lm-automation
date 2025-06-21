package main

import (
	"os"
	"strings"

	"github.com/endgate-systems/lm-automation/cmd"
)

func main() {
	// Bypass full Listmonk init if running CLI tool
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "test-") {
		cmd.Execute()
		return
	}

	// Default fallback: (optional - can be removed)
	cmd.Execute()
}
