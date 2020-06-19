package handlers

import (
	"djale/pkg/entities"
	"djale/pkg/store"
	"encoding/json"
	"github.com/kellydunn/golang-geo"
	"log"
	"net/http"
)

func (a *App) GetVanIDForDestination(w http.ResponseWriter, r *http.Request) {

	var request entities.VanRunRequest
	vanStore := store.NewVanStore(a.DB)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vans, err := vanStore.GetAllVans()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Unable to get vans")

	}

	vanID := 0
	valid := false

	p := geo.NewPoint(request.Latitude, request.Longitude)

	// this will give us the first vanID which will have 500km radius of requested destination and destination that goes to
	for _, v := range vans {
		p2 := geo.NewPoint(v.Latitude, v.Longitude)
		dist := p.GreatCircleDistance(p2)
		if dist < 500 {
			valid = true
			vanID = v.ID
		}

		if valid {
			RespondWithJSON(w, http.StatusOK, vanID)
		}

	}

}
