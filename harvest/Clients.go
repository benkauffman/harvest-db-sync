package harvest

import (
	"encoding/json"
	"log"
	"fmt"
	"github.com/benkauffman/harvest-db-sync/models"
)

func GetAllClients() ([]models.Client) {

	var list []models.Client
	var page, count int
	count = 1

	for page = 1; page <= count; page++ {

		fmt.Println(fmt.Sprintf("requesting clients page [%d] from harvest", page))

		wrapper := models.HarvestClientWrapper{}
		jsonErr := json.Unmarshal(requestClients(page), &wrapper)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		count = wrapper.TotalPages

		if page == 1 {
			list = wrapper.Clients
		} else {
			for i := range wrapper.Clients {
				single := models.Client{}
				single = wrapper.Clients[i]
				list = append(list, single)
			}
		}
	}

	fmt.Println(fmt.Sprintf("[%d] clients retrieved from harvest", len(list)))

	return list
}
