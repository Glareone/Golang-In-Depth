package database

import (
	"log"
)

func migrationsAlterTables() {
	// this migration commented because otherwise I need to prepare the whole mechanism
	// or use GORM to track migrations in the separate table
	// Jan_11_2025_AlterEventsTable()
}

func Jan_11_2025_AlterEventsTable() {
	transaction, err := DB.Begin()

	if err != nil {
		log.Fatalf("[Alter Table] Transaction creation failed", err)
	}

	defer func() {
		if err != nil {
			rbErr := transaction.Rollback()
			if rbErr != nil {
				log.Fatalf("[Rollback Error] Failed to rollback transaction: %v", rbErr)
			}
			log.Printf("[Update Error] Rolled back due to failure: %v", err)
		}
	}()

	_, err = transaction.Exec("UPDATE events SET user_id = NULL WHERE user_id IS NOT NULL")
	if err != nil {
		transaction.Rollback()
		log.Fatalf("[UPDATE EVENTS] Events within events table cannot be updated, failed: %w", err)
	}

	alterEventsAddReferenceToUsersTableQuery :=
		`
    ALTER TABLE events
    ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id) 
    REFERENCES users(id)
    ON DELETE SET NULL;
	`

	_, err = transaction.Exec(alterEventsAddReferenceToUsersTableQuery)
	if err != nil {
		transaction.Rollback()
		log.Fatalf("[Alter Table] Events table update failed: %w", err)
	}

	err = transaction.Commit() // commit the transaction
	if err != nil {
		log.Fatalf("[Transaction] Transaction Commit Failed: %w", err)
	}
}
