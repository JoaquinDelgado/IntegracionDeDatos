package Handlers

import (
	"IntegracionDeDatos/Internal/Services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBeachesByUserProfileID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var idAux int64
	idAux, _ = strconv.ParseInt(id, 10, 32)
	beaches, _ := Services.GetBeachesByUserProfileID(int(idAux))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beaches)
}
