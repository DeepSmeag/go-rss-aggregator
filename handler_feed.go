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

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	//TODO: validation
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		log.Printf("Error 400 on /feeds: %v", err)
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, "Could not create feed")
		log.Printf("Error 400 on /feeds: %v", err)
		return
	}

	respondWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, "Could not get feeds")
		log.Printf("Error 400 on /feeds: %v", err)
		return
	}

	respondWithJSON(w, 200, models.DatabaseFeedsToFeeds(feeds))
}
