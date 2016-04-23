package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/evcraddock/supp-store/config"
	"github.com/evcraddock/supp-store/routers"
	//"github.com/gorilla/mux"
)

func main() {

	configuration := config.LoadFile()
	websiteAddress := configuration.Website.Address + ":" + configuration.Website.Port

	router := routers.InitRoutes()
	n := negroni.Classic()

	fmt.Printf("start service on port %v \n", websiteAddress)

	n.UseHandler(&router)
	log.Fatal(http.ListenAndServe(websiteAddress, n))

}

// type WithCORS struct {
// 	r *mux.Router
// }

// func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	if origin := req.Header.Get("Origin"); origin != "" {
// 		res.Header().Set("Access-Control-Allow-Origin", origin)
// 		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		res.Header().Set("Access-Control-Allow-Headers",
// 			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 	}

// 	//.Printf("%v", origin)
// 	// Stop here for a Preflighted OPTIONS request.
// 	if req.Method == "OPTIONS" {
// 		fmt.Printf("handled options")
// 		return
// 	}
// 	// Lets Gorilla work
// 	s.r.ServeHTTP(res, req)
// }
