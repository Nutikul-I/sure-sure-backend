package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/textures1245/payso-check-slip-backend/util"
)

// Replace with your own connection parameters
var server = "x"
var port = 5432
var user = "x"
var password = "x"
var database = "x"
var DB *sql.DB

func Init() {
	server = viper.GetString("DEV_DB_HOST")
	port = 5432
	user = viper.GetString("DEV_DB_USER")
	password = viper.GetString("DEV_DB_PASSWORD")
	database = viper.GetString("DEV_DB_NAME")
	// Create connection string
	var err error

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)

	// Create connection pool
	DB, err = sql.Open("postgres", connString)

	if err != nil {
		log.Error("**** Error creating connection pool: " + err.Error())
	}
	log.Debug("==-- Connected! --==")
	log.Infof("connString %s", connString)
}

func ConnectDB() *sql.DB {
	var err error

	util.Init()
	Init()
	if DB == nil {
		connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)
		log.Infof("connString %s", connString)
		DB, err = sql.Open("postgres", connString)
		if err != nil {
			log.Error("**** Error creating connection pool: " + err.Error())
		}
	}

	log.Debug("==-- Connected! --==")
	return DB
}
