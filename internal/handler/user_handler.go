package handler

import (
	"errors"
	"github.com/AriSu2904/go-auth/internal/service"
	"github.com/AriSu2904/go-auth/internal/utils"
	"net/http"
)

type UserHandler interface {
	FindByPersona(w http.ResponseWriter, r *http.Request)
	FindByEmail(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) FindByPersona(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	persona := query.Get("persona")

	if persona == "" {
		utils.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST_BODY",
			"Persona cannot be empty")
		return
	}

	user, err := h.userService.FindByPersona(r.Context(), &persona)

	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, "USER_NOT_FOUND",
				"User with the given persona does not exist")
			return
		} else {
			utils.WriteError(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR",
				"An unexpected error occurred")
			return
		}
	}

	response := map[string]interface{}{
		"message": "Successfully retrieved user",
		"data":    user,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *userHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {

}
