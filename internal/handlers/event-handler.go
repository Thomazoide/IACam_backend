package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Thomazoide/IACam_backend/internal/db"
	"github.com/Thomazoide/IACam_backend/internal/models"
	"github.com/Thomazoide/IACam_backend/internal/utils"
	"github.com/Thomazoide/IACam_backend/internal/ws"
	"github.com/go-chi/chi/v5"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	DB := db.GetInstance()
	var event *models.Event
	encoder := json.NewEncoder(w)
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Fatal(err.Error())
		encoder.Encode(utils.ResponseWriter(err.Error(), nil, true))
		return
	}
	event.Checked = false
	DB.Create(&event)
	encoder.Encode(utils.ResponseWriter("Evento guardado", event, false))
}

func CheckEvent(w http.ResponseWriter, r *http.Request) {
	DB := db.GetInstance()
	var event *models.Event
	id := chi.URLParam(r, "id")
	encoder := json.NewEncoder(w)
	if err := DB.First(&event, id).Error; err != nil {
		log.Fatal(err.Error())
		encoder.Encode(utils.ResponseWriter(err.Error(), nil, true))
		return
	}
	event.Checked = true
	if err := DB.Save(&event).Error; err != nil {
		log.Fatal(err.Error())
		encoder.Encode(utils.ResponseWriter(err.Error(), nil, true))
		return
	}
	encoder.Encode(utils.ResponseWriter("Evento actualizado", event, false))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	DB := db.GetInstance()
	id := chi.URLParam(r, "id")
	encoder := json.NewEncoder(w)
	var event *models.Event
	if err := DB.First(&event, id).Error; err != nil {
		log.Fatal(err.Error())
		encoder.Encode(utils.ResponseWriter(err.Error(), nil, true))
		return
	}
	if err := DB.Delete(&event).Error; err != nil {
		log.Fatal(err.Error())
	}
}

func CatchEvent(w http.ResponseWriter, r *http.Request) {
	var e *models.Event
	json.NewDecoder(r.Body).Decode(&e)
	msg, err := json.Marshal(e)
	if err != nil {
		ws.Broadcast([]byte(err.Error()))
		return
	}
	ws.Broadcast(msg)
	w.WriteHeader(http.StatusOK)
}
