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

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responder.WithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing json", err))
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responder.WithError(w, http.StatusBadRequest, fmt.Sprint("Error creating user", err))
		return
	}

	responder.WithJson(w, http.StatusOK, DatabaseUserToUser(user))
}
