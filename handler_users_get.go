package main

import (
	"net/http"

	"github.com/grodier/bootdev-rss-aggregator/internal/database"
	"github.com/grodier/bootdev-rss-aggregator/internal/responder"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	responder.WithJson(w, http.StatusOK, DatabaseUserToUser(user))
}
