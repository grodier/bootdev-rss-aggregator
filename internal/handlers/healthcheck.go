package handlers

import (
	"net/http"

	"github.com/grodier/bootdev-rss-aggregator/internal/responder"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {

	type healthcheckResponse struct {
		Status string `json:"status"`
	}

	responder.WithJson(w, http.StatusOK, healthcheckResponse{
		Status: "ok",
	})
}
