package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Thomazoide/IACam_backend/internal/payloads"
	"github.com/go-chi/chi/v5"
)

func StreamProxy(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	id := chi.URLParam(r, "id")
	workerURL := "http://worker-camera-" + id + ":5000/stream"
	resp, err := http.Get(workerURL)
	if err != nil {
		encoder.Encode(&payloads.ResponsePayload{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		})
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	io.Copy(w, resp.Body)
}
