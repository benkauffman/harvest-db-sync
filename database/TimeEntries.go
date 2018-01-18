package database

import (
	"../models"
	"strings"
	"log"
)

func SaveTimeEntries(list []models.TimeEntry) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.

	collection := sess.Collection("time_entry")
	for i, element := range list {

		_, err := collection.Insert(element)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Printf("time entry list index %d update time entry %d", i, element.ID)
				res := collection.Find("id", element.ID)
				var found models.TimeEntry
				res.One(&found)
				res.Update(element)
			} else {
				log.Printf("time entry list index %d insert time entry %d", i, element.ID)
				log.Printf("db.Insert(): %q\n", err)
			}
		} else {
			log.Printf("time entry list index %d insert time entry %d", i, element.ID)
		}
	}
}
