package main

import (
	"log"
	"net/http"

	"github.com/Thomazoide/IACam_backend/internal/config"
	"github.com/Thomazoide/IACam_backend/internal/db"
	"github.com/Thomazoide/IACam_backend/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := config.Load()
	db.Connect(*cfg)
	router := routes.SetupRouter()
	http.ListenAndServe(":8080", router)
	log.Println("Servidor en el puerto :8080")
}
