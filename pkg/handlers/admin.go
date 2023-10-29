package handlers

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/forms"
	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
)

func (m *Repository) Admin(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "admin.page.html", &models.TemplateData{})
}

func (m *Repository) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}
