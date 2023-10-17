package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ecoarchie/go-rss/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title string `json:"title"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json, %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title: params.Title,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Couldn't create feed, %v", err))
		return
	}

	respondWithJSON(w, 201, dbFeedToFeed(feed))
}

// func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request, user database.User) {

// 	respondWithJSON(w, 200, dbUserToUser(user))
// }

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Couldn't get feed, %v", err))
		return
	}

	respondWithJSON(w, 200, dbFeedsToFeeds(feeds))
}