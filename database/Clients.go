package database

import (
	"../models"
	"strings"
	"log"
)

func SaveClients(list []models.Client) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.

	collection := sess.Collection("client")
	for i, element := range list {

		_, err := collection.Insert(element)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Printf("client list index %d update client %d", i, element.ID)
				res := collection.Find("id", element.ID)
				var found models.Client
				res.One(&found)
				res.Update(element)
			} else {
				log.Printf("client list index %d insert client %d", i, element.ID)
				log.Printf("db.Insert(): %q\n", err)
			}
		} else {
			log.Printf("client list index %d insert client %d", i, element.ID)
		}
	}
}
