package handler

import (
	"errors"
	"github.com/AriSu2904/go-auth/internal/service"
	"github.com/AriSu2904/go-auth/internal/utils"
	"log"
	"net/http"
)

type UserHandler interface {
	FindByQuery(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) FindByQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	persona := query.Get("persona")
	email := query.Get("email")
	log.Print("Incoming request with query params - persona: ", persona, ", email: ", email)

	if len(persona) > 0 {
		userPersona, err := h.userService.FindByPersona(r.Context(), &persona)

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
			"data":    userPersona,
		}
		utils.WriteJSON(w, http.StatusOK, response)
		return
	}

	if len(email) > 0 {
		userEmail, err := h.userService.FindByEmail(r.Context(), &email)

		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				utils.WriteError(w, http.StatusNotFound, "USER_NOT_FOUND",
					"User with the given email does not exist")
				return
			} else {
				utils.WriteError(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR",
					"An unexpected error occurred")
				return
			}
		}

		response := map[string]interface{}{
			"message": "Successfully retrieved user",
			"data":    userEmail,
		}
		utils.WriteJSON(w, http.StatusOK, response)
		return
	}

	utils.WriteError(w, http.StatusBadRequest, "BAD_REQUEST", "Query parameter 'persona' or 'email' is required")
}
