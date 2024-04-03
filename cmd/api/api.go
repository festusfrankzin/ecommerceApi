package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/festusfrankzin/ecommerceApi/services/user"
	"github.com/gorilla/mux"
)

// Creating a struct that takes the address and database
type ApiServer struct { 
	address string
	db      *sql.DB

}


// Create a new ApiServer
func CreateApiServer(address string , db * sql.DB) *ApiServer {

	return &ApiServer{
		address: address,
		db: db,
	}

}
// Running the server
func (s *ApiServer) Run ()error{

	// Creating a router using gorilla mux 
	router := mux.NewRouter()


	// Creating a subrouter for api version 1
	subRouter := router.PathPrefix("/api/v1/").Subrouter()
	

	userHandler := user.Newhandler()
	userHandler.RegisterRoutes(subRouter)


	log.Println("Listening on ", s.address)
	return http.ListenAndServe(s.address, router)

}