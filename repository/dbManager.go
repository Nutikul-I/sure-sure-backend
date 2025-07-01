package repository

import (
	"database/sql"
	"fmt"
	"time"

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
	// อ่าน config เฉพาะครั้งแรก
	if server == "x" {
		server = viper.GetString("DEV_DB_HOST")
		port = 5432
		user = viper.GetString("DEV_DB_USER")
		password = viper.GetString("DEV_DB_PASSWORD")
		database = viper.GetString("DEV_DB_NAME")
	}

	// สร้าง connection เฉพาะเมื่อ DB ยัง nil
	if DB == nil {
		var err error
		connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)

		// Create connection pool
		DB, err = sql.Open("postgres", connString)

		if err != nil {
			log.Error("**** Error creating connection pool: " + err.Error())
		} else {
			// ตั้งค่า connection pool เพื่อป้องกัน connection leak
			DB.SetMaxOpenConns(10)                  // จำกัด connection สูงสุด 10 ตัว
			DB.SetMaxIdleConns(5)                   // จำกัด idle connection 5 ตัว
			DB.SetConnMaxLifetime(30 * time.Minute) // connection อายุสูงสุด 30 นาที
		}

		log.Debug("==-- Connected! --==")
		log.Infof("connString %s", connString)
	}
}

func ConnectDB() *sql.DB {
	// ใช้ singleton pattern - สร้าง connection เพียงครั้งเดียว
	if DB == nil {
		util.Init()
		Init()
	}

	log.Debug("==-- Connected! --==")
	return DB
}
