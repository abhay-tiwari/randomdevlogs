package handlers

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
)

func (m *Repository) Algorithms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "algorithms.page.html", &models.TemplateData{})
}

func (m *Repository) BubbleSort(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "bubble-sort.page.html", &models.TemplateData{})
}

func (m *Repository) LinearSearch(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "linear-search.page.html", &models.TemplateData{})
}
