package main

import (
	"fmt"
	"net/http"

	"github.com/grodier/bootdev-rss-aggregator/internal/auth"
	"github.com/grodier/bootdev-rss-aggregator/internal/responder"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responder.WithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v\n", err))
		return
	}

	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		responder.WithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v\n", err))
		return
	}

	responder.WithJson(w, http.StatusOK, DatabaseUserToUser(user))
}
