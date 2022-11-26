package Handlers

import (
	"IntegracionDeDatos/Internal/Services"
	"encoding/json"
	"net/http"
)

func InsertBusLines(w http.ResponseWriter, r *http.Request) {
	Services.InsertBusLines()
}

func InsertRowsInLifeGuardStationsBusStops(w http.ResponseWriter, r *http.Request) {
	err := Services.InsertRowsInLifeGuardStationsBusStops()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
}
