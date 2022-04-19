package api

import (
	"IbanValidator/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Result   bool
	ErrorExp string
}

func NewHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/validateIban/{iban}", validateHandler).Methods("POST")

	return r
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iban := vars["iban"]

	if !validataInput(iban) {
		http.Error(w, "iban null", 400)
		return
	}

	err, validateIbanResponse := service.ValidateIban(iban)

	w.Header().Set("Content-Type", "application/json")
	if err {
		response := Response{Result: false, ErrorExp: validateIbanResponse}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	} else {
		response := Response{Result: true}
		json.NewEncoder(w).Encode(response)
	}

}
