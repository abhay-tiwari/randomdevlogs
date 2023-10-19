package helpers

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
)

var app *config.AppConfig

func IsAuthenticated(r *http.Request) bool {
	tokenExists := app.Session.Exists(r.Context(), "user_id")

	return tokenExists
}
