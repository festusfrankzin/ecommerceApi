package user

import (
	"github.com/festusfrankzin/ecommerceApi/models"
	"github.com/festusfrankzin/ecommerceApi/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func Newhandler() *Handler {

	return &Handler{}

}

//Registering Routes

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleLogin).Methods("POST")
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// Loads the Json data from the server
	// Check if the person is already registered
	// If not, Create a new profile


	var payload models.RegisterUserPayload
	if err := utils.ParseJson(r, payload); err != nil {

		utils.WriteError(w, http.StatusBadRequest, err)
	}


}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}
