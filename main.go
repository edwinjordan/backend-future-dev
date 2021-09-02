package main

import (
	"log"

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

	dbUser := "sql6433460"
	dbPass := "6ZueX9pJFG"
	dbHost := "localhost"
	dbName := "sql6433460"
	dbPort := "3306"

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
