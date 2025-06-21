package automation

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func StartScheduler(db *sqlx.DB, logger *log.Logger) {
	logger.Println("Automation scheduler started...")

	go func() {
		for {
			// TODO: Fetch trigger queue or pending actions.
			// MVP: Just log and sleep to prove it's running.
			logger.Println("Checking for triggers...")

			time.Sleep(30 * time.Second)
		}
	}()
}
