package harvest

import (
	"encoding/json"
	"log"
	"fmt"
	"../models"
)

func GetAllTimeEntries() ([]models.TimeEntry) {

	var list []models.TimeEntry
	var page, count int
	count = 1

	for page = 1; page <= count; page++ {

		fmt.Println(fmt.Sprintf("requesting time entries page [%d] from harvest", page))

		wrapper := models.HarvestTimeEntryWrapper{}
		jsonErr := json.Unmarshal(requestTimeEntries(page), &wrapper)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		count = wrapper.TotalPages

		if page == 1 {
			list = wrapper.TimeEntries
		} else {
			for i := range wrapper.TimeEntries {
				single := models.TimeEntry{}
				single = wrapper.TimeEntries[i]
				list = append(list, single)
			}
		}
	}

	fmt.Println(fmt.Sprintf("[%d] time entries retrieved from harvest", len(list)))

	return list
}
