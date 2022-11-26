package Handlers

import (
	"IntegracionDeDatos/Internal/Services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBusLines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var idAux int64
	idAux, _ = strconv.ParseInt(id, 10, 32)
	busLines, err := Services.GetBusLinesByBeachID(idAux)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(busLines)
}
