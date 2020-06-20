package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) initializeRoutes() {

	//customer handlers
	a.Router.HandleFunc("/customer", a.CreateConsumer).Methods("POST")
	a.Router.HandleFunc("/customer/{id:[0-9]+}", a.GetCustomerByID).Methods("GET")
	a.Router.HandleFunc("/customer/{id:[0-9]+}", a.UpdateCustomerByID).Methods("PUT")
	a.Router.HandleFunc("/customer/{id:[0-9]+}", a.DeleteCustomerByID).Methods("DELETE")

	//consignment handlers
	a.Router.HandleFunc("/consignment", a.CreateConsignment).Methods("POST")
	a.Router.HandleFunc("/consignment/{id:[0-9]+}", a.GetConsignmentByID).Methods("GET")
	a.Router.HandleFunc("/consignment/{id:[0-9]+}", a.UpdateConsignmentByID).Methods("PUT")
	a.Router.HandleFunc("/consignment/{id:[0-9]+}", a.DeleteConsignmentByID).Methods("DELETE")
	a.Router.HandleFunc("/consignments", a.InsertConsignmentsFromCSV).Methods("POST")

	//van handlers
	a.Router.HandleFunc("/van", a.GetVanIDForDestination).Methods("POST")

	//order handlers
	a.Router.HandleFunc("/order/{order_id:[0-9]+}", a.GetVansForOrderID).Methods("POST")



	// Swagger
	a.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
