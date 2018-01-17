package database

import (
	"time"
	"strconv"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"github.com/benkauffman/harvest-db-sync/config"
)

func UpdateSyncMillis() {
	UpdateAppData("sync.millis", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
}

func GetSyncMillis() (int64) {
	var millis, _ = strconv.ParseInt(GetAppData("sync.millis"), 10, 64)
	return millis
}

func GetSyncDate() (string) {

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

	var date = configuration.Sync.DefaultDate

	var millis = GetSyncMillis()

	if millis > 1 {
		var instance = time.Unix(0, millis*int64(time.Millisecond)).AddDate(0, 0, -1)
		date = fmt.Sprintf("%d-%d-%d", instance.Year(), instance.Month(), instance.Day())
	}

	return date
}
