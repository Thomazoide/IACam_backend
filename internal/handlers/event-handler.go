package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Thomazoide/IACam_backend/internal/ws"
)

type Event struct {
	CameraID string `json:"camera_id"`
	Event    string `json:"event"`
}

func CatchEvent(w http.ResponseWriter, r *http.Request) {
	var e Event
	json.NewDecoder(r.Body).Decode(&e)
	msg, err := json.Marshal(e)
	if err != nil {
		ws.Broadcast([]byte(err.Error()))
		return
	}
	ws.Broadcast(msg)
	w.WriteHeader(http.StatusOK)
}
