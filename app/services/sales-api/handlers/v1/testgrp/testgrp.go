package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/PhyoYazar/service/foundation/web"
)

func Test(ctx context.Context,w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errors.New("UNTRUSTED ERROR")
	 }

	// 1. Validate the data
	// 2. Call into the business layer
	// 3. Return errors
	// 4. Handle OK response

	status := struct{
		Status string `json:"status"`
	}{
		Status: "Ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
