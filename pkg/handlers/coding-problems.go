package handlers

import (
	"log"
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
)

func (m *Repository) GetCodingProblemsPage(w http.ResponseWriter, r *http.Request) {
	category := "coding-problems"
	blogs, err := m.DB.GetBlogsByCategory(category)

	if err != nil {
		log.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["blogs"] = blogs

	var td models.TemplateData
	td.Data = data

	render.RenderTemplate(w, r, "coding-problems.page.html", &td)
}
