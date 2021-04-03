package connector

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() *sql.DB {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_NAME")
	dbConf := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName+"?parseTime=true&&loc=Asia%2FBangkok"

	db, _ := sql.Open("mysql", dbConf)
	_, err := db.Exec("SET NAMES utf8")
	if err != nil {
		log.Fatal(err)
	}
	_, errTimeZone := db.Exec("SET GLOBAL time_zone = '+7:00'")
	if errTimeZone != nil {
		log.Fatal(err)
	}

	return db
}
