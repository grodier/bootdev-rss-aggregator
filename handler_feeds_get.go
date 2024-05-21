package main

import (
	"fmt"
	"net/http"

	"github.com/grodier/bootdev-rss-aggregator/internal/responder"
)

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		responder.WithError(w, http.StatusBadRequest, fmt.Sprint("Error getting feeds", err))
		return
	}

	responder.WithJson(w, http.StatusCreated, DatabaseFeedsToFeeds(feeds))
}
