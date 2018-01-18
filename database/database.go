package database

import (
	"log"
	"upper.io/db.v3/mysql" // Imports the mysql adapter.
	"upper.io/db.v3/lib/sqlbuilder"
	"github.com/spf13/viper"
	"../config"
)

func getDatabase() (sqlbuilder.Database) {

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

	conf := configuration.Database

	var settings = mysql.ConnectionURL{
		Database: conf.Database,
		Host:     conf.Host,
		User:     conf.User,
		Password: conf.Password,
	}

	// Attempting to establish a connection to the database.
	sess, err := mysql.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	return sess

}
