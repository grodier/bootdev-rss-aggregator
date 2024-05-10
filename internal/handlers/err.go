package handlers

import (
	"net/http"

	"github.com/grodier/bootdev-rss-aggregator/internal/responder"
)

func Err(w http.ResponseWriter, r *http.Request) {
	responder.WithError(w, http.StatusInternalServerError, "Internal Server Error")
}
