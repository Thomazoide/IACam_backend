package utils

import "github.com/Thomazoide/IACam_backend/internal/payloads"

func ResponseWriter(message string, data any, err bool) *payloads.ResponsePayload {
	return &payloads.ResponsePayload{
		Message: message,
		Data:    data,
		Error:   err,
	}
}
