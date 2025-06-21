package automation

import (
	"fmt"
	"time"

	"internal/automation/engine"
	"internal/hooks"
	"github.com/knadh/listmonk/models"
)

func init() {
	engine.RegisterRule(engine.Rule{
		Name:        "welcome-flow-1",
		Trigger:     "subscriber_created",
		Description: "Send a welcome email 30 minutes after subscriber joins List X",
		Condition: func(sub models.Subscriber) bool {
			targetListUUID := "REPLACE_WITH_LIST_UUID" // Replace
			for _, l := range sub.Lists {
				if l.UUID == targetListUUID {
					return true
				}
			}
			return false
		},
		Action: func(sub models.Subscriber) {
			time.AfterFunc(30*time.Minute, func() {
				templateID := 123 // Replace

				err := hooks.SendTemplateEmail(sub.ID, templateID)
				if err != nil {
					fmt.Printf("❌ Failed to send email to %s: %v\n", sub.Email, err)
				} else {
					fmt.Printf("✅ Email sent to %s using template %d\n", sub.Email, templateID)
				}
			})
		},
	})
}
