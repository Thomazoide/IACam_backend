package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Thomazoide/IACam_backend/internal/db"
	"github.com/Thomazoide/IACam_backend/internal/models"
	"github.com/Thomazoide/IACam_backend/internal/payloads"
	"github.com/Thomazoide/IACam_backend/internal/services"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateCamera(w http.ResponseWriter, r *http.Request) {
	var DB *gorm.DB = db.GetInstance()
	var cam *models.Camera
	encoder := json.NewEncoder(w)
	if err := json.NewDecoder(r.Body).Decode(&cam); err != nil {
		log.Fatal(err.Error())
		encoder.Encode(&payloads.ResponsePayload{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		})
		return
	}
	cam.Status = "active"
	DB.Create(&cam)
	go services.CreateWorker(cam.ID, cam.RTSP)
	encoder.Encode(&payloads.ResponsePayload{
		Message: "Cámara agregada",
		Data:    cam,
		Error:   false,
	})
}

func GetCameras(w http.ResponseWriter, r *http.Request) {
	var DB *gorm.DB = db.GetInstance()
	var cams []models.Camera
	encoder := json.NewEncoder(w)
	if err := DB.Find(&cams).Error; err != nil {
		encoder.Encode(&payloads.ResponsePayload{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		})
		return
	}
	encoder.Encode(&payloads.ResponsePayload{
		Message: "Cámaras registradas",
		Data:    cams,
		Error:   false,
	})
}

func GetCamera(w http.ResponseWriter, r *http.Request) {
	var DB *gorm.DB = db.GetInstance()
	id := chi.URLParam(r, "id")
	encoder := json.NewEncoder(w)
	var cam *models.Camera
	if err := DB.First(&cam, id).Error; err != nil {
		encoder.Encode(&payloads.ResponsePayload{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		})
		return
	}
	encoder.Encode(&payloads.ResponsePayload{
		Message: "Cámara encontrada",
		Data:    cam,
		Error:   false,
	})
}

func DeleteCamera(w http.ResponseWriter, r *http.Request) {
	var DB *gorm.DB = db.GetInstance()
	id := chi.URLParam(r, "id")
	encoder := json.NewEncoder(w)
	var cam *models.Camera
	if err := DB.First(&cam, id).Error; err != nil {
		encoder.Encode(&payloads.ResponsePayload{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		})
	}
	go services.RemoveWorker(cam.ID)
	if err := DB.Delete(&cam).Error; err != nil {
		encoder.Encode(&payloads.ResponsePayload{
			Message: err.Error(),
			Data:    nil,
			Error:   true,
		})
		return
	}
	encoder.Encode(&payloads.ResponsePayload{
		Message: "Cámara eliminada",
		Data:    nil,
		Error:   false,
	})
}
