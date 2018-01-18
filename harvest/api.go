package harvest

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
	"strconv"
	"../config"
	"../database"
)

func requestClients(page int) ([]byte) {
	return get("clients", page, fmt.Sprintf("&updated_since=%s", database.GetSyncDate()))
}

func requestInvoices(page int) ([]byte) {
	return get("invoices", page, fmt.Sprintf("&updated_since=%s", database.GetSyncDate()))
}

func requestProjects(page int) ([]byte) {
	return get("projects", page, fmt.Sprintf("&updated_since=%s", database.GetSyncDate()))
}

func requestTasks(page int) ([]byte) {
	return get("tasks", page, fmt.Sprintf("&updated_since=%s", database.GetSyncDate()))
}

func requestTimeEntries(page int) ([]byte) {
	return get("time_entries", page, fmt.Sprintf("&updated_since=%s", database.GetSyncDate()))
}

func requestUsers(page int) ([]byte) {
	return get("users", page, fmt.Sprintf("&updated_since=%s", database.GetSyncDate()))
}

func get(endPoint string, page int, params string) ([]byte) {

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	conf := configuration.Harvest

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(conf.BaseUrl+endPoint+"?page=%d"+params, page), nil)
	req.Header.Set("User-Agent", conf.AppName)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.Token))
	req.Header.Set("Harvest-Account-ID", strconv.Itoa(conf.AccountId))

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
