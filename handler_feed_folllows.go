package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rohan556/rss-generator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		handleError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		handleError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		handleError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feeds))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId := chi.URLParam(r, "feed_follow_id")

	feedFollowUuid, err := uuid.Parse(feedFollowId)

	if err != nil {
		handleError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowUuid,
		UserID: user.ID,
	})

	if err != nil {
		handleError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	respondWithJSON(w, 200, "Unfollowed successfully")
}
