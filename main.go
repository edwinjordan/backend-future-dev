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
	dbUser, dbPassword, dbName, dbHost := "sql6433460", "6ZueX9pJFG", "sql6433460", "localhost:3306"

	db, err := database.ConnectToDB(dbUser, dbPassword, dbName, dbHost)

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

	route.Run(":8000")

}
