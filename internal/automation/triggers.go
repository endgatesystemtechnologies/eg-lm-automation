package automation

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type AutomationTrigger struct {
	ID         int       `db:"id"`
	Type       string    `db:"type"`
	TargetID   int       `db:"target_id"`
	Action     string    `db:"action"`
	ActionData string    `db:"action_data"`
	Delay      int       `db:"delay"`
	CreatedAt  time.Time `db:"created_at"`
}

// FetchTriggers fetches all automation triggers from the DB.
func FetchTriggers(db *sqlx.DB, logger *log.Logger) ([]AutomationTrigger, error) {
	query := `
	SELECT id, type, target_id, action, action_data, delay, created_at
	FROM automation_triggers
	`
	var triggers []AutomationTrigger
	err := db.Select(&triggers, query)
	if err != nil {
		logger.Println("Error fetching triggers:", err)
		return nil, err
	}
	return triggers, nil
}
