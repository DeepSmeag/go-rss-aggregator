package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DeepSmeag/go-rss-aggregator/api/models"
	"github.com/DeepSmeag/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	//TODO: validation, there needs to be a name check (I don't think the video handles this)
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		log.Printf("Error 400 on /users: %v", err)
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, "Could not create user")
		log.Printf("Error 400 on /users: %v", err)
		return
	}

	respondWithJSON(w, 200, models.DatabaseUserToUser(user))
}
