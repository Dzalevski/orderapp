package handlers

import (
	"djale/pkg/entities"
	"djale/pkg/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


// GetCustomerByID godoc
// @Summary Get details of customer by id
// @Description Get details of customer by id
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {object} Customer
// @Router /customer/{id} [get]
func (a *App) GetCustomerByID(w http.ResponseWriter, r *http.Request) {

	customer := &entities.Customers{}

	custStore := store.NewCustomerStore(a.DB)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	customer, err = custStore.GetCustomerByID(id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Customer not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, customer)
}

func (a *App) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	var c entities.Customers
	custStore := store.NewCustomerStore(a.DB)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := custStore.InsertCustomer(c.Name, c.Address, c.Postcode, c.GeoLocation); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to create customer")
		return
	}

	RespondWithJSON(w, http.StatusCreated, "Customer successfully created")
}

func (a *App) UpdateCustomerByID(w http.ResponseWriter, r *http.Request) {
	custStore := store.NewCustomerStore(a.DB)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var cust entities.Customers
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cust); err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := custStore.UpdateCustomerByID(cust.Name, cust.Address, cust.Postcode, cust.GeoLocation, id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to update customer by id")
		return
	}

	RespondWithJSON(w, http.StatusOK, "Customer successfully updated")
}

func (a *App) DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	custStore := store.NewCustomerStore(a.DB)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
		return
	}

	if err := custStore.DeleteCustomerByID(id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Successfully deleted customer")
}
