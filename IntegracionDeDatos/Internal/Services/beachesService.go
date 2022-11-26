package Services

import (
	"IntegracionDeDatos/Domain"
	"IntegracionDeDatos/Repositories"
)

func GetBeachesByUserProfileID(userProfileID int) ([]Domain.Beach, error) {
	beaches, err := Repositories.BeachesByUserProfileID(userProfileID)
	if err != nil {
		return nil, err
	}
	return beaches, nil
}
