package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Thomazoide/IACam_backend/internal/payloads"
	"github.com/go-chi/chi/v5"
)

func StreamProxy(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	id := chi.URLParam(r, "id")
	workerURL := "http://worker-camera-" + id + ":5000/stream"
	client := &http.Client{
		Timeout: 0,
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			ResponseHeaderTimeout: 5 * time.Second,
		},
	}
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, workerURL, nil)
	if err != nil {
		encoder.Encode(&payloads.ResponsePayload{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		})
		return
	}
	resp, err := client.Do(req)
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
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(resp.StatusCode)

	flusher, _ := w.(http.Flusher)
	buf := make([]byte, 32*1024)
	for {
		n, readErr := resp.Body.Read(buf)
		if n > 0 {
			if _, writeErr := w.Write(buf[:n]); writeErr != nil {
				return
			}
			if flusher != nil {
				flusher.Flush()
			}
		}
		if readErr != nil {
			if readErr == io.EOF {
				return
			}
			return
		}
	}
}
