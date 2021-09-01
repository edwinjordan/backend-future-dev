package services

import (
	"go-heroku/dtos"
	"go-heroku/models"
	"go-heroku/repositories"
)

func FindOneWilayahById(link string, repository repositories.WilayahRepository) dtos.Response {
	operationResult := repository.FindOneById(link)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Wilayah)

	return dtos.Response{Success: true, Data: data}
}
