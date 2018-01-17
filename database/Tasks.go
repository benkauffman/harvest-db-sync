package database

import (
	"github.com/benkauffman/harvest-db-sync/models"
	"strings"
	"log"
)

func SaveTasks(list []models.Task) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.

	collection := sess.Collection("task")
	for i, element := range list {

		_, err := collection.Insert(element)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Printf("task list index %d update task %d", i, element.ID)
				res := collection.Find("id", element.ID)
				var found models.Task
				res.One(&found)
				res.Update(element)
			} else {
				log.Printf("task list index %d insert task %d", i, element.ID)
				log.Printf("db.Insert(): %q\n", err)
			}
		} else {
			log.Printf("task list index %d insert task %d", i, element.ID)
		}
	}
}
