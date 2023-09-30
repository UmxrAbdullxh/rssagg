package main

import (
	"net/http"
)

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	type successResponse struct {
		Status string `json:"status"`
	}
	responseWithJson(w, 200, successResponse{
		Status: "ok",
	})
}
