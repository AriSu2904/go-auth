package handler

import (
	"encoding/json"
	"errors"
	"github.com/AriSu2904/go-auth/internal/dto"
	"github.com/AriSu2904/go-auth/internal/service"
	"github.com/AriSu2904/go-auth/internal/utils"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authService service.AuthService
	logger      *slog.Logger
	validate    *validator.Validate
}

func NewAuthHandler(authService service.AuthService, log *slog.Logger) AuthHandler {
	return &authHandler{authService: authService, logger: log, validate: validator.New()}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Process incoming registration request", "layer", "authHandler")

	var requestBody dto.RegisterUserInput

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST_BODY",
			"Invalid request body make sure you have all the required fields")
		return
	}

	err = h.validate.Struct(requestBody)
	if err != nil {
		utils.WriteValidationError(w, err)
		return
	}

	user, err := h.authService.SignUp(r.Context(), &requestBody)

	if err != nil {
		if errors.Is(err, service.ErrUserExists) {
			utils.WriteError(w, http.StatusConflict, "USER_ALREADY_EXISTS",
				service.ErrUserExists.Error())
			return
		} else {
			utils.WriteError(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR",
				"An unexpected error occurred")
			return
		}
	}

	response := map[string]interface{}{
		"message": "Registration successful",
		"data": map[string]string{
			"email":   user.Email,
			"persona": user.Persona,
		},
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Process incoming login request", "layer", "authHandler")

	var payload dto.LoginUserInput

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "BAD_REQUEST",
			"Invalid request body make sure you have email/persona and password")
		return
	}

	err = h.validate.Struct(payload)
	if err != nil {
		utils.WriteValidationError(w, err)
		return
	}

	tokenInfo, err := h.authService.SignIn(r.Context(), &payload, nil)

	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			utils.WriteError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS",
				"Email/Persona or Password is incorrect")
			return
		} else {
			utils.WriteError(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR",
				"An unexpected error occurred")
			return
		}
	}

	response := map[string]interface{}{
		"message": "Login successful",
		"data":    tokenInfo,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
