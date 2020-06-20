package handlers

import (
	"djale/pkg/entities"
	"djale/pkg/service"
	"djale/pkg/store"
	"encoding/json"
	"log"
	"net/http"
)

func (a *App) GetVanIDForDestination(w http.ResponseWriter, r *http.Request) {

	var request entities.VanRunRequest
	destService := service.NewDestinationService(store.NewVanStore(a.DB))

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
