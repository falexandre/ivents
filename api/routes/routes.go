package routes

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/peterwade153/ivents/api/controllers"
	"github.com/peterwade153/ivents/api/middlewares"
	"github.com/peterwade153/ivents/api/responses"
)

func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Ivents")
}

// Handlers routes
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/register", controllers.UserSignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	s := r.PathPrefix("/api").Subrouter()
	s.Use(middlewares.AuthJwtVerify)
	s.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	s.HandleFunc("/venues", controllers.CreateVenue).Methods("POST")
	s.HandleFunc("/venues", controllers.GetVenues).Methods("GET")
	s.HandleFunc("/venues/{id}", controllers.UpdateVenue).Methods("PUT")
	s.HandleFunc("/venues/{id}", controllers.DeleteVenue).Methods("DELETE")

	return r
}
