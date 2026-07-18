package repositories

import (
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/client/models"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/database"
)

func FindByEmail(email string) (*models.Client, error) {

	var client models.Client

	err := database.DB.
		Where("email = ?", email).
		First(&client).Error

	if err != nil {
		return nil, err
	}

	return &client, nil

}
