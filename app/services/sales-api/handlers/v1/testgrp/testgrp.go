package testgrp

import (
	"context"
	"encoding/json"
	"net/http"
)

func Test(ctx context.Context,w http.ResponseWriter, r *http.Request) error {

	// 1. Validate the data
	// 2. Call into the business layer
	// 3. Return errors
	// 4. Handle OK response

	status := struct{
		Status string `json:"status"`
	}{
		Status: "Ok",
	}

	return json.NewEncoder(w).Encode(status)
}
