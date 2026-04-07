package db

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
	"APPOINMENT_BOOKING_SYSTEM/config"
)

var (
	DB *sql.DB
	err error
)

func ConnectMySQL() {
	cfg := mysql.Config{
		User:                 config.AppConfig.MySQL.User,
		Passwd:               config.AppConfig.MySQL.Password,
		Net:                  config.AppConfig.MySQL.Net,
		Addr:                 config.AppConfig.MySQL.Address,
		DBName:               config.AppConfig.MySQL.DBName,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("MySQL connection failed: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("MySQL ping failed: %v", err)
	}

	DB = db
	fmt.Println("MySQL connected successfully")
}

func CloseMySQL() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing MySQL: %v", err)
		} else {
			fmt.Println("MySQL connection closed.")
		}
	}
}
