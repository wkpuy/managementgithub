package main

import (
	"fmt"
	"os"

	"management/internal/http/route"
	"management/internal/utilities/path"

	"github.com/joho/godotenv"
)

func main() {
	rootDir := path.GetRootDirectory()

	// load .env for development
	godotenv.Load(rootDir + "/.env")
	os.Setenv("TZ", "Asia/Bangkok")

	// connect mysql
	// mysqlDB := connector.ConnectMysql()
	// defer func() {
	// 	err := mysqlDB.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("Connection to MySQL closed.")
	// }()

	 fmt.Println("Run")

	// serve http request
	serverPort := os.Getenv("SERVER_PORT")
	route.HandleAPI(serverPort, rootDir)
}
