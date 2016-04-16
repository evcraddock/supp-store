package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/evcraddock/supp-store/config"
	"github.com/evcraddock/supp-store/routers"
)

func main() {

	configuration := config.LoadFile()
	websiteAddress := configuration.Website.Address + ":" + configuration.Website.Port

	router := routers.InitRoutes()
	n := negroni.Classic()

	fmt.Printf("start service on port %v \n", websiteAddress)
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(websiteAddress, n))

}
