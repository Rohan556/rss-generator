package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rohan556/rss-generator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `name`
		URL  string `url`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		handleError(w, 400, fmt.Sprint("Error with parsing JSON", err))
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
		handleError(w, 400, fmt.Sprintf("Couldn't create user %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())

	if err != nil {
		handleError(w, 400, "SOmething went wrong")
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
