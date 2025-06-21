package hooks

import (
    "fmt"

    "github.com/knadh/listmonk/models"
    "github.com/endgate-systems/lm-automation/internal/automation"
)

func OnSubscriberCreated(sub models.Subscriber) {
    fmt.Printf("🔥 Hook: Subscriber created - %s\n", sub.Email)
    automation.RunAutomations("subscriber_created", sub)
}
