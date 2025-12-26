package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/shivasharanappaBiradar2645/tinker-lab/golang/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
	//"github.com/go-chi/cors"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	dbstring := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Fatal("cant connect to db: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()
	//router.Use(cors.Handler(cors.Options{}))

	v1Router := chi.NewRouter()
	//v1Router.HandleFunc("/health", handlerReadiness)

	v1Router.Get("/health", handlerReadiness)
	v1Router.HandleFunc("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server starting on port %v", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
