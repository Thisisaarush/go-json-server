package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/thisisaarush/go-rss/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	// PORT
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	type apiConfig struct {
		DB *database.Queries
	}
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL is not set")
	}

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	apiCfg := &apiConfig{	DB: database.New(conn) }

	

 // Router
	router := chi.NewRouter()

	// Cors
	router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, 
  }))
	
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	log.Printf("Server started on port %s", PORT)
	err = http.ListenAndServe(":" + PORT, router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}