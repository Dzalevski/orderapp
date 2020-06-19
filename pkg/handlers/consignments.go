package handlers

import (
	"djale/pkg/entities"
	"djale/pkg/store"
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)


func (a *App) GetConsignmentByID(w http.ResponseWriter, r *http.Request) {

	cons := &entities.Consignment{}

	consStore := store.NewConsignmentsStore(a.DB)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid consignments ID")
		return
	}

	cons, err = consStore.GetConsignmentByID(id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Consignments not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, cons)
}

func (a *App) CreateConsignment(w http.ResponseWriter, r *http.Request) {
	var c entities.Consignment
	consStore := store.NewConsignmentsStore(a.DB)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := consStore.InsertConsignment(c.Barcode, c.LinkToSupplier); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to create consignments")
		return
	}

	RespondWithJSON(w, http.StatusCreated, "Consignments successfully created")
}

func (a *App) UpdateConsignmentByID(w http.ResponseWriter, r *http.Request) {
	consStore := store.NewConsignmentsStore(a.DB)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid consignments ID")
		return
	}

	var cons entities.Consignment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cons); err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := consStore.UpdateConsignmentByID(cons.Barcode, cons.LinkToSupplier, id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to update consignments by id")
		return
	}

	RespondWithJSON(w, http.StatusOK, "Consignments successfully updated")
}

func (a *App) DeleteConsignmentByID(w http.ResponseWriter, r *http.Request) {
	consStore := store.NewConsignmentsStore(a.DB)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Consignments ID")
		return
	}

	if err := consStore.DeleteConsignmentByID(id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Successfully deleted consignments")
}

func (a *App) InsertConsignmentsFromCSV(w http.ResponseWriter, r *http.Request) {

	consStore := store.NewConsignmentsStore(a.DB)

	err := r.ParseMultipartForm(10 << 22)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	file, _, err := r.FormFile("data")
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	if len(headers) < 1 {
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		err = consStore.InsertConsignment(record[0], record[1])
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Unable to create consignments")
			return
		}

	}
	RespondWithJSON(w, http.StatusOK, "Consignments successfully created")

}
