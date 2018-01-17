package harvest

import (
	"encoding/json"
	"log"
	"fmt"
	"github.com/benkauffman/harvest-db-sync/models"
)

func GetAllTasks() ([]models.Task) {

	var list []models.Task
	var page, count int
	count = 1

	for page = 1; page <= count; page++ {

		fmt.Println(fmt.Sprintf("requesting tasks page [%d] from harvest", page))

		wrapper := models.HarvestTaskWrapper{}
		jsonErr := json.Unmarshal(requestTasks(page), &wrapper)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		count = wrapper.TotalPages

		if page == 1 {
			list = wrapper.Tasks
		} else {
			for i := range wrapper.Tasks {
				single := models.Task{}
				single = wrapper.Tasks[i]
				list = append(list, single)
			}
		}
	}

	fmt.Println(fmt.Sprintf("[%d] tasks retrieved from harvest", len(list)))

	return list
}
