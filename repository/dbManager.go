package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/textures1245/payso-check-slip-backend/util"
)

// Replace with your own connection parameters
var server = "x"
var port = 1433
var user = "x"
var password = "x"
var database = "x"
var DB *sql.DB

func Init() {
	server = viper.GetString("DEV_DB_HOST")
	port = 1433
	user = viper.GetString("DEV_DB_USER")
	password = viper.GetString("DEV_DB_PASSWORD")
	database = viper.GetString("DEV_DB_NAME")
	// Create connection string
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)

	// Create connection pool
	DB, err = sql.Open("sqlserver", connString)

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
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)
		log.Infof("connString %s", connString)
		DB, err = sql.Open("sqlserver", connString)
		if err != nil {
			log.Error("**** Error creating connection pool: " + err.Error())
		}
	}

	log.Debug("==-- Connected! --==")
	return DB
}
