package database

import (
	"log"
)

func UpdateAppData(attribute string, value string) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.

	_, err := sess.Exec(`INSERT INTO app_data (attribute, value) 
	VALUES (?,?) ON DUPLICATE KEY UPDATE value=VALUES(value)`, attribute, value)

	if err != nil {
		log.Printf("db.Insert(): %q\n", err)
	}
}

func GetAppData(attribute string) (string) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.
	row, err := sess.QueryRow(`SELECT value FROM app_data WHERE attribute = ? LIMIT ?`, attribute, 1)

	if err != nil {
		log.Printf("db.Select (): %q\n", err)
	}

	var value string
	row.Scan(&value)
	return value
}
