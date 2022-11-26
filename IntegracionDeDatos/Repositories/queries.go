package Repositories

import (
	"IntegracionDeDatos/Domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func BeachesByUserProfileID(id int) ([]Domain.Beach, error) {
	// An albums slice to hold data from returned rows.
	var beaches []Domain.Beach

	db := getDBInstance()
	rows, err := db.Query("SELECT id, name, address, coordinates, flag_colour, chemical_bathroom, has_market FROM lifeguardstations WHERE id in (SELECT lifeguardstations_id FROM lifeguardstations_user_profile WHERE user_profile_id = ?) ", id)
	if err != nil {
		return nil, fmt.Errorf("beachesByUserProfile %v: %v", id, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var beachDAO Domain.BeachDAO
		if err := rows.Scan(&beachDAO.ID, &beachDAO.Name, &beachDAO.Address, &beachDAO.Coordinates, &beachDAO.FlagColor, &beachDAO.ChemicalBathroom, &beachDAO.HasMarket); err != nil {
			return nil, fmt.Errorf("beachesByUserProfile %v: %v", id, err)
		}
		var ints []float64
		err := json.Unmarshal([]byte(beachDAO.Coordinates), &ints)
		if err != nil {
			log.Fatal(err)
		}
		var beach = Domain.Beach{ID: beachDAO.ID, Name: beachDAO.Name, Address: beachDAO.Address, Coordinates: ints, FlagColor: beachDAO.FlagColor, ChemicalBathroom: beachDAO.ChemicalBathroom, HasMarket: beachDAO.HasMarket}
		beaches = append(beaches, beach)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("beachesByUserProfile %v: %v", id, err)
	}
	return beaches, nil
}

func GetBeaches() ([]Domain.Beach, error) {
	// An albums slice to hold data from returned rows.
	var beaches []Domain.Beach

	db := getDBInstance()
	rows, err := db.Query("SELECT * FROM lifeguardstations")
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la query GetBeaches, el error es: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var beachDAO Domain.BeachDAO
		if err := rows.Scan(&beachDAO.ID, &beachDAO.Name, &beachDAO.Address, &beachDAO.Coordinates, &beachDAO.FlagColor, &beachDAO.ChemicalBathroom, &beachDAO.HasMarket); err != nil {
			return nil, fmt.Errorf("error al mapear la fila a la estructura, el error es: %v", err)
		}
		var ints []float64
		err := json.Unmarshal([]byte(beachDAO.Coordinates), &ints)
		if err != nil {
			log.Fatal(err)
		}
		var beach = Domain.Beach{ID: beachDAO.ID, Name: beachDAO.Name, Address: beachDAO.Address, Coordinates: ints, FlagColor: beachDAO.FlagColor, ChemicalBathroom: beachDAO.ChemicalBathroom, HasMarket: beachDAO.HasMarket}

		beaches = append(beaches, beach)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getBeaches : %v", err)
	}
	return beaches, nil
}

func GetBusStopIDs() ([]Domain.BusStopID, error) {
	var busStopIDs []Domain.BusStopID

	db := getDBInstance()
	rows, err := db.Query("SELECT busstopId FROM busstops where buslines is null")
	if err != nil {
		return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var busStopID Domain.BusStopID
		if err := rows.Scan(&busStopID.ID); err != nil {
			return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
		}
		busStopIDs = append(busStopIDs, busStopID)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
	}
	return busStopIDs, nil
}

func GetBusStops() ([]Domain.BusStop, error) {
	var busStops []Domain.BusStop

	db := getDBInstance()
	rows, err := db.Query("SELECT * FROM busstops")
	if err != nil {
		return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var busStopDAO Domain.BusStopDAO
		if err := rows.Scan(&busStopDAO.ID, &busStopDAO.Street1, &busStopDAO.Street2, &busStopDAO.Street1ID, &busStopDAO.Street2ID, &busStopDAO.Location, &busStopDAO.BusLines); err != nil {
			return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
		}
		var location Domain.Location
		err := json.Unmarshal([]byte(busStopDAO.Location), &location)
		if err != nil {
			log.Fatal(err)
		}
		var busStop = Domain.BusStop{ID: busStopDAO.ID, Street1: busStopDAO.Street1, Street2: busStopDAO.Street2, Street1ID: busStopDAO.Street1ID, Street2ID: busStopDAO.Street2ID, Location: location, BusLines: busStopDAO.BusLines}

		busStops = append(busStops, busStop)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
	}
	return busStops, nil
}

func GetLifeGuardStationsBusLinesByBeachID(ID int64) ([]Domain.BusLine, error) {
	var busLines []Domain.BusLine

	db := getDBInstance()
	rows, err := db.Query("SELECT buslines FROM lifeguardstationsbusstops where lifeguardstationId = ?", ID)
	if err != nil {
		return nil, fmt.Errorf("GetLifeGuardStationsBusLinesByBeachID error was: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var aux []byte
		var aux2 []Domain.BusLine
		if err := rows.Scan(&aux); err != nil {
			return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
		}
		err := json.Unmarshal(aux, &aux2)
		if err != nil {
			return nil, err
		}
		for _, Line := range aux2 {
			busLines = append(busLines, Line)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBusStopIDs error was: %v", err)
	}
	return busLines, nil

}

//backoffice queries

func InsertBusLines(busStopID int, busLines []Domain.BusLine) {
	db := getDBInstance()
	tx, _ := db.Begin()
	cmd := "UPDATE busstops SET buslines = ? WHERE busstopId = ?"
	updateBusLines, errPrepare := tx.Prepare(cmd)
	if errPrepare != nil {
		panic(errPrepare)
	}
	bl, _ := json.Marshal(busLines)
	_, errExec := updateBusLines.Exec(bl, busStopID)
	if errExec != nil {
		panic(errExec)
	}
	updateBusLines.Close()
	errCommit := tx.Commit()
	if errCommit != nil {
		panic(errCommit)
	}
}

func InsertRowsInLifeGuardStationsBusStops(beach Domain.Beach, busStop Domain.BusStop) error {
	db := getDBInstance()
	query := "INSERT INTO lifeguardstationsbusstops(busstopId, street1, street2, street1Id, street2Id, location, buslines, lifeguardstationId, name, address, coordinates, flag_colour, chemical_bathroom, has_market) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	loc, _ := json.Marshal(busStop.Location)
	cord, _ := json.Marshal(beach.Coordinates)
	res, err := stmt.ExecContext(ctx, busStop.ID, busStop.Street1, busStop.Street2, busStop.Street1ID, busStop.Street2ID, loc, busStop.BusLines, beach.ID, beach.Name, beach.Address, cord, beach.FlagColor, beach.ChemicalBathroom, beach.HasMarket)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d inserts created ", rows)
	return nil
}
