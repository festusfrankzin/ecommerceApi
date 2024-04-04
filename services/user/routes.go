package user 

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Handler struct {

}

func Newhandler() *Handler {
	
	return &Handler{

	}

}

//Registering Routes 

func (h *Handler)RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.HandleLogin).Methods("POST")
	router.HandleFunc("/register",h.HandleLogin).Methods("POST")
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request){
	 
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request){
	
}