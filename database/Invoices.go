package database

import (
	"github.com/benkauffman/harvest-db-sync/models"
	"strings"
	"log"
)

func SaveInvoices(list []models.Invoice) {
	sess := getDatabase()
	defer sess.Close() // Remember to close the database session.

	collection := sess.Collection("invoice")
	for i, element := range list {

		_, err := collection.Insert(element)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Printf("invoice list index %d update invoice %d", i, element.ID)
				res := collection.Find("id", element.ID)
				var found models.Invoice
				res.One(&found)
				res.Update(element)
			} else {
				log.Printf("invoice list index %d insert invoice %d", i, element.ID)
				log.Printf("db.Insert(): %q\n", err)
			}
		} else {
			log.Printf("invoice list index %d insert invoice %d", i, element.ID)
		}
	}
}
