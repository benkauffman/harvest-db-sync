package harvest

import (
	"encoding/json"
	"log"
	"fmt"
	"../models"
)

func GetAllUsers() ([]models.User) {

	var list []models.User
	var page, count int
	count = 1

	for page = 1; page <= count; page++ {

		fmt.Println(fmt.Sprintf("requesting users page [%d] from harvest", page))

		wrapper := models.HarvestUserWrapper{}
		jsonErr := json.Unmarshal(requestUsers(page), &wrapper)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		count = wrapper.TotalPages

		if page == 1 {
			list = wrapper.Users
		} else {
			for i := range wrapper.Users {
				single := models.User{}
				single = wrapper.Users[i]
				list = append(list, single)
			}
		}
	}

	fmt.Println(fmt.Sprintf("[%d] users retrieved from harvest", len(list)))

	return list
}
