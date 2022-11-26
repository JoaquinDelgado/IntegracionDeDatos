package Services

import (
	"IntegracionDeDatos/External/Services"
	"IntegracionDeDatos/Repositories"
	"fmt"
	"github.com/umahmood/haversine"
	"time"
)

func InsertBusLines() {
	busStopIDs, err := Repositories.GetBusStopIDs()
	if err != nil {
		panic(err)
	}
	for _, busStopID := range busStopIDs {
	INTENTARDEVUELTA:
		fmt.Printf("el bus stop id es: %v \n", busStopID.ID)
		busLines, errGet := Services.GetBusLines(busStopID.ID)
		if errGet != nil {
			fmt.Printf("el error fue :%v \n", errGet)
			time.Sleep(60 * time.Second)
			goto INTENTARDEVUELTA
		}
		Repositories.InsertBusLines(busStopID.ID, busLines)
	}
}

func InsertRowsInLifeGuardStationsBusStops() error {
	// obtener playas
	beaches, _ := Repositories.GetBeaches()
	// obtener busstops
	busStops, _ := Repositories.GetBusStops()
	// for playa in playas
	for _, beach := range beaches {
		for _, busStop := range busStops {
			coordenadaBeach := haversine.Coord{Lat: beach.Coordinates[1], Lon: beach.Coordinates[0]}                         // Oxford, UK
			coordenadaBusStop := haversine.Coord{Lat: busStop.Location.Coordinates[1], Lon: busStop.Location.Coordinates[0]} // Turin, Italy
			_, km := haversine.Distance(coordenadaBeach, coordenadaBusStop)
			if km <= 0.5 {
				err := Repositories.InsertRowsInLifeGuardStationsBusStops(beach, busStop)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
