package harvest

import (
	"encoding/json"
	"log"
	"fmt"
	"../models"
)

func GetAllInvoices() ([]models.Invoice) {

	var list []models.Invoice
	var page, count int
	count = 1

	for page = 1; page <= count; page++ {

		fmt.Println(fmt.Sprintf("requesting invoices page [%d] from harvest", page))

		wrapper := models.HarvestInvoiceWrapper{}
		jsonErr := json.Unmarshal(requestInvoices(page), &wrapper)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		count = wrapper.TotalPages

		if page == 1 {
			list = wrapper.Invoices
		} else {
			for i := range wrapper.Invoices {
				single := models.Invoice{}
				single = wrapper.Invoices[i]
				list = append(list, single)
			}
		}
	}

	fmt.Println(fmt.Sprintf("[%d] invoices retrieved from harvest", len(list)))

	return list
}
