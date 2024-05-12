package main

import (
	"net/http"

	"github.com/grodier/bootdev-rss-aggregator/internal/responder"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responder.WithError(w, http.StatusInternalServerError, "Internal Server Error")
}
