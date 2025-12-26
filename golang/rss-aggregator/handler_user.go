package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shivasharanappaBiradar2645/tinker-lab/golang/rss-aggregator/internal/auth"
	"github.com/shivasharanappaBiradar2645/tinker-lab/golang/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decoder.Decode(&params)

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	log.Print(err)
	log.Print(params.Name)

	respondWithJSON(w, 201, user)
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("%v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("%v", err))
		return
	}

	respondWithJSON(w, 200, user)

}
