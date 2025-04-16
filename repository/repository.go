package repository

import (
	"database/sql"
	"log"
)

// RpMsg represents a record in the rp_msg table.
type RpMsg struct {
	ID        int
	Body      string
	CreateTime string
}

// GetAllRpMsgs retrieves all records from the rp_msg table.
func GetAllRpMsgs(db *sql.DB) ([]RpMsg, error) {
	query := "SELECT id, body, create_time FROM rp_msg"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []RpMsg
	for rows.Next() {
		var msg RpMsg
		if err := rows.Scan(&msg.ID, &msg.Body, &msg.CreateTime); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
