package main

import (
	"djale/pkg/handlers"
)

// @title Orders API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	//app run
	a := handlers.App{}
	a.Initialize("postgres", "djale12345", "postgres")
	a.Run(":8080")

}
