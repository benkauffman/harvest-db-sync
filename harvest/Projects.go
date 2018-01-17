package harvest

import (
	"encoding/json"
	"log"
	"fmt"
	"github.com/benkauffman/harvest-db-sync/models"
)

func GetAllProjects() ([]models.Project) {

	var list []models.Project
	var page, count int
	count = 1

	for page = 1; page <= count; page++ {

		fmt.Println(fmt.Sprintf("requesting projects page [%d] from harvest", page))

		wrapper := models.HarvestProjectWrapper{}
		jsonErr := json.Unmarshal(requestProjects(page), &wrapper)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		count = wrapper.TotalPages

		if page == 1 {
			list = wrapper.Projects
		} else {
			for i := range wrapper.Projects {
				single := models.Project{}
				single = wrapper.Projects[i]
				list = append(list, single)
			}
		}
	}

	fmt.Println(fmt.Sprintf("[%d] projects retrieved from harvest", len(list)))

	return list
}
