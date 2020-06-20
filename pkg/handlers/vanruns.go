package handlers

import (
	"djale/pkg/entities"
	"djale/pkg/service"
	"djale/pkg/store"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (a *App) GetVanIDForDestination(w http.ResponseWriter, r *http.Request) {

	var request entities.VanRunRequest
	destService := service.NewDestinationService(store.NewVanStore(a.DB), store.NewOrderStore(a.DB), store.NewCustomerStore(a.DB))

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vanID, err := destService.GetDestinationBetweenTwoPoints(request.Latitude, request.Longitude)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Unable to calculate distance")
		return
	}
	RespondWithJSON(w, http.StatusOK, vanID)

}

// this one have hackish solution because of short time
func (a *App) GetVansForOrderID(w http.ResponseWriter, r *http.Request) {
	destService := service.NewDestinationService(store.NewVanStore(a.DB), store.NewOrderStore(a.DB), store.NewCustomerStore(a.DB))

	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["order_id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	resp, err := destService.WhichVanForCustomerOrder(orderID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "No vans available")
		return
	}

	file, err := os.Create("result.csv")
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Unable to process the request")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	var s []string
	vIdsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(resp.VanIDs)), ","), "[]")

	s = append(s, vIdsString, string(resp.ConsID))
	err = writer.Write(s)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Unable to process the request")
		return
	}

}
