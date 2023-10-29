package handlers

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/forms"
	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
)

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())

	form := forms.New(r.PostForm)

	form.Required("email", "password")
	form.IsEmail("email", r)

	if !form.Valid() {
		data := make(map[string]interface{})

		data["user"] = models.User{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		render.RenderTemplate(w, r, "login.page.html", &models.TemplateData{
			Data: data,
			Form: form,
		})

		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	id, _, err := m.DB.Authenticate(email, password)

	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Invalid Login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	m.App.Session.Put(r.Context(), "user_id", id)

	m.App.Session.Put(r.Context(), "flash", "Logged in Successfully.")

	http.Redirect(w, r, "/admin/all-blogs", http.StatusSeeOther)
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
