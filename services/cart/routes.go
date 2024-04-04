package cart

import (
	"net/http"
	"github.com/gorilla/mux"
)


type Handler struct {

}

func Newhandler()*Handler { 

	return &Handler{

	}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/cart", h.cartHandler).Methods(http.MethodGet)
}

func (h *Handler) cartHandler(w http.ResponseWriter, r *http.Request){

}