package services

import (
	"errors"

	"github.com/rusbandtecnologia/rusband-api-mobile/internal/client/models"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/client/repositories"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/utils"
)

func Login(email string, password string) (*models.Client, error) {

	client, err := repositories.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(client.Password, password) {
		return nil, errors.New("senha inválida")
	}

	return client, nil

}
