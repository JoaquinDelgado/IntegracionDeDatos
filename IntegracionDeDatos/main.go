package main

import (
	"IntegracionDeDatos/Handlers"
	"IntegracionDeDatos/Repositories"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	Repositories.InitMYSQLConnection()
	r := mux.NewRouter()
	r.HandleFunc("/buslines/beach/{id}/buses", Handlers.GetBusLines).Methods("GET")
	r.HandleFunc("/lifeguardstations/userprofile/{id}/beaches", Handlers.GetBeachesByUserProfileID).Methods("GET")

	//backoffice
	r.HandleFunc("/backoffice/insertBusLines", Handlers.InsertBusLines).Methods("GET")
	r.HandleFunc("/backoffice/InsertRowsInLifeGuardStationsBusStops", Handlers.InsertRowsInLifeGuardStationsBusStops).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
