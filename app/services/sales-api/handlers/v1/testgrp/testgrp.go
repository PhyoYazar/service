package testgrp

import (
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	// 1. Validate the data
	// 2. Call into the business layer
	// 3. Return errors
	// 4. Handle OK response

	status := struct{
		Status string `json:"status"`
	}{
		Status: "Ok",
	}

	json.NewEncoder(w).Encode(status)
}
