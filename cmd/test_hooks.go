package cmd

import (
    "fmt"

    "github.com/endgate-systems/lm-automation/internal/hooks"
    "github.com/knadh/listmonk/models"
    "github.com/spf13/cobra"
)

var testHooksCmd = &cobra.Command{
    Use:   "test-hooks",
    Short: "Simulate subscriber_created event and trigger automations",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(">>> TEST-HOOKS COMMAND TRIGGERED <<<")
        sub := models.Subscriber{
            UUID:  "test-uuid",
            Email: "test@example.com",
            Name:  "Test User",
        }
        hooks.OnSubscriberCreated(sub)
    },
}

func init() {
    rootCmd.AddCommand(testHooksCmd)
}
