package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/grodier/bootdev-rss-aggregator/internal/database"
	"github.com/grodier/bootdev-rss-aggregator/internal/responder"
)

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responder.WithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing json", err))
		return
	}

	if params.Name == "" || params.URL == "" {
		responder.WithError(w, http.StatusBadRequest, fmt.Sprint("Missing params in body"))
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		responder.WithError(w, http.StatusBadRequest, fmt.Sprint("Error creating feed", err))
		return
	}

	responder.WithJson(w, http.StatusCreated, DatabaseFeedToFeed(feed))
}
