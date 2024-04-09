package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/PhyoYazar/service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/PhyoYazar/service/business/web/auth"
	"github.com/PhyoYazar/service/business/web/v1/mid"
	"github.com/PhyoYazar/service/foundation/web"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Auth 		*auth.Auth
}

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) * web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test )
	app.Handle(http.MethodGet, "/test/auth", testgrp.Test , mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAdminOnly))

	return app
}