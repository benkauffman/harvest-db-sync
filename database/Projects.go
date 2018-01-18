package database

import (
	"../models"
	"strings"
	"log"
)

func SaveProjects(list []models.Project) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.

	collection := sess.Collection("project")
	for i, element := range list {

		_, err := collection.Insert(element)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Printf("project list index %d update project %d", i, element.ID)
				res := collection.Find("id", element.ID)
				var found models.Project
				res.One(&found)
				res.Update(element)
			} else {
				log.Printf("project list index %d insert project %d", i, element.ID)
				log.Printf("db.Insert(): %q\n", err)
			}
		} else {
			log.Printf("project list index %d insert project %d", i, element.ID)
		}
	}
}
