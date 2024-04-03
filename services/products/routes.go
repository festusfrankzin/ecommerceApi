package products

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func Newhandler() *Handler {

	return &Handler{}

}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/products", h.handlerProducts).Methods(http.MethodGet)
	router.HandleFunc("/products:id", h.EachhandlerProducts).Methods(http.MethodGet)

}


func (h *Handler) handlerProducts(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) EachhandlerProducts(w http.ResponseWriter, r *http.Request){

}