package middleware

import (
	"github.com/AriSu2904/go-auth/internal/dto"
	"github.com/AriSu2904/go-auth/internal/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func HeaderValidator(next http.Handler) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deviceId := r.Header.Get("X-Device-ID")
		deviceInfo := r.Header.Get("X-Device-Info")

		addsHeader := dto.AdditionalHeader{
			DeviceId:   deviceId,
			DeviceInfo: deviceInfo,
		}

		err := validate.Struct(addsHeader)

		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST_HEADER",
				"X-Device-ID and X-Device-Info headers are required")
			return
		}

		next.ServeHTTP(w, r)
	})
}
