package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"abc_pharmacy/database"
)

func RegisterInvoiceRoutes(router *mux.Router) {
	router.HandleFunc("/api/invoices", getInvoices).Methods("GET")
	router.HandleFunc("/api/invoices", createInvoice).Methods("POST")
	// Add additional routes for editing and deleting invoices
}

func getInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := database.GetInvoices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func createInvoice(w http.ResponseWriter, r *http.Request) {
	var newInvoice database.Invoice
	err := json.NewDecoder(r.Body).Decode(&newInvoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.AddNewInvoice(newInvoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

