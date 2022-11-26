package Services

import (
	"IntegracionDeDatos/Domain"
	"IntegracionDeDatos/Repositories"
)

func GetBusLinesByBeachID(ID int64) ([]Domain.BusLine, error) {
	busLines, err := Repositories.GetLifeGuardStationsBusLinesByBeachID(ID)
	if err != nil {
		return nil, err
	}
	//eliminar repetidos
	inResult := make(map[Domain.BusLine]bool)
	var result []Domain.BusLine
	for _, busLine := range busLines {
		inMap := inResult[busLine]
		if !inMap {
			inResult[busLine] = true
			result = append(result, busLine)
		}
	}
	return result, nil
}
