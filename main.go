package main

import (
	"log"
	"os"

	"go-heroku/configs"
	"go-heroku/database"
	"go-heroku/models"
	"go-heroku/repositories"
)

func main() {
	//database configs
	// errEnv := godotenv.Load()

	// if errEnv != nil {
	// 	panic("Failed to load env file")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	db, err := database.ConnectToDB(dbUser, dbPass, dbHost, dbName, dbPort)

	if err != nil {
		log.Fatalln(err)
	}

	//ping to database
	err = db.DB().Ping()

	//error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	//migration
	db.AutoMigrate(&models.Wilayah{})

	defer db.Close()

	wilayahRepository := repositories.NewWilayahRepository(db)

	route := configs.SetupRoutes(wilayahRepository)

	route.Run()

}
