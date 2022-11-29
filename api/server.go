package api

import (
	"fmt"
	"log"
	"os"

	"go-morclinic/api/controllers"
	"go-morclinic/api/seed"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// LIVE
	// server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	// DEV
	server.Initialize(os.Getenv("TestDbDriver"), os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbPort"), os.Getenv("TestDbHost"), os.Getenv("TestDbName"))

	seed.Load(server.DB)

	server.Run(":8080")

}
