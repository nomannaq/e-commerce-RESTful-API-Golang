package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/services/auth"
	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/types"
	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleLogin).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.RegisterUserPayload

	if err := utils.ParsejSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//Check if user exists

	_, err := h.store.GetUsersByEmail(payload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//Create user if it doesn't exist
	_, err = h.store.CreateUser(&types.User{
		Email:     payload.Email,
		Password:  hashedPassword,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, "user created successfully")

}
