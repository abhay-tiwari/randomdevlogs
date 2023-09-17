package handlers

import (
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.page.html")
}
