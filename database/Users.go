package database

import (
	"github.com/benkauffman/harvest-db-sync/models"
	"strings"
	"log"
)

func SaveUsers(list []models.User) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.

	collection := sess.Collection("user")
	for i, element := range list {

		_, err := collection.Insert(element)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Printf("user list index %d update user %d", i, element.ID)
				res := collection.Find("id", element.ID)
				var found models.User
				res.One(&found)
				res.Update(element)
			} else {
				log.Printf("user list index %d insert user %d", i, element.ID)
				log.Printf("db.Insert(): %q\n", err)
			}
		} else {
			log.Printf("user list index %d insert user %d", i, element.ID)
		}
	}
}
