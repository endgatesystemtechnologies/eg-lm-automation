package automation

import (
	"time"
)

type Trigger struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	TriggerType  string    `db:"trigger_type"`
	TriggerValue string    `db:"trigger_value"`
	ActionType   string    `db:"action_type"`
	ActionValue  string    `db:"action_value"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
