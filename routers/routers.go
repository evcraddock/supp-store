package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() CORSHandler {
	router := mux.NewRouter()
	router = SetSupplementEntryRoutes(router)

	return CORSHandler{router}
}

type CORSHandler struct {
	r *mux.Router
}

func (s *CORSHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if req.Method == "OPTIONS" {
		return
	}

	s.r.ServeHTTP(res, req)
}
