package main

import (
	"./harvest"
	"./database"
	"log"
)

func main() {
	log.Printf("starting sync using %s as start date", database.GetSyncDate())

	database.SaveClients(harvest.GetAllClients())
	database.SaveUsers(harvest.GetAllUsers())
	database.SaveProjects(harvest.GetAllProjects())
	database.SaveInvoices(harvest.GetAllInvoices())
	database.SaveTasks(harvest.GetAllTasks())
	database.SaveTimeEntries(harvest.GetAllTimeEntries())

	database.UpdateSyncMillis()
}
