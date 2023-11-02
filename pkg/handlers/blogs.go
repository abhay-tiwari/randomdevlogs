package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
	"github.com/go-chi/chi/v5"
	"github.com/gomarkdown/markdown"
)

func (m *Repository) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, _ := m.DB.GetAllBlogs()

	data := make(map[string]interface{})

	data["blogs"] = blogs

	data["totalBlogs"] = len(blogs)

	var templateData models.TemplateData

	templateData.Data = data

	render.RenderTemplate(w, r, "all-blogs.page.html", &templateData)
}

func (m *Repository) AddBlog(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "add-blog.page.html", &models.TemplateData{})
}

func (m *Repository) EditBlog(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "blogId")

	id, err := strconv.Atoi(blogId)

	if err != nil {
		log.Println(err)
		return
	}

	blog, err := m.DB.GetBlogById(id)

	if err != nil {
		log.Println("err is in fetch blog")
		log.Println(err)
		return
	}

	data := make(map[string]interface{})

	data["blog"] = blog

	var td models.TemplateData
	td.Data = data

	render.RenderTemplate(w, r, "edit-blog.page.html", &td)
}

func (m *Repository) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := r.Form.Get("blogId")
	blogId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		return
	}

	m.DB.DeleteBlog(blogId)

	http.Redirect(w, r, "/admin/all-blogs", http.StatusSeeOther)
}

func (m *Repository) PostEditBlog(w http.ResponseWriter, r *http.Request) {
	title := r.Form.Get("title")
	metaDescription := r.Form.Get("metaDescription")
	ogTitle := r.Form.Get("ogTitle")
	ogDescription := r.Form.Get("ogDescription")
	slug := r.Form.Get("slug")
	category := r.Form.Get("category")
	tags := r.Form.Get("tags")
	createdBy := r.Form.Get("createdBy")

	file, _, err := r.FormFile("content")

	if err != nil {
		return
	}

	defer file.Close()

	fileContent, err := io.ReadAll(file)

	html := markdown.ToHTML(fileContent, nil, nil)

	blog := models.Blog{
		Title:           title,
		MetaDescription: metaDescription,
		OgTitle:         ogTitle,
		OgDescription:   ogDescription,
		Slug:            slug,
		Category:        category,
		Tags:            tags,
		CreatedBy:       createdBy,
		Content:         template.HTML(html),
	}

	m.DB.EditBlog(blog)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func (m *Repository) SubmitBlog(w http.ResponseWriter, r *http.Request) {
	title := r.Form.Get("title")
	metaDescription := r.Form.Get("metaDescription")
	ogTitle := r.Form.Get("ogTitle")
	ogDescription := r.Form.Get("ogDescription")
	slug := r.Form.Get("slug")
	category := r.Form.Get("category")
	tags := r.Form.Get("tags")
	createdBy := r.Form.Get("createdBy")

	file, _, err := r.FormFile("content")

	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()

	fileContent, err := io.ReadAll(file)

	html := markdown.ToHTML(fileContent, nil, nil)

	blog := models.Blog{
		Title:           title,
		MetaDescription: metaDescription,
		OgTitle:         ogTitle,
		OgDescription:   ogDescription,
		Slug:            slug,
		Category:        category,
		Tags:            tags,
		CreatedBy:       createdBy,
		Content:         template.HTML(html),
	}

	err = m.DB.AddBlog(blog)
	if err != nil {
		log.Println(err)
		return
	}

	m.App.Session.Put(r.Context(), "flash", "Blog Added Successfully.")

	http.Redirect(w, r, "/admin/all-blogs", http.StatusSeeOther)
}

func (m *Repository) GetBlogPage(w http.ResponseWriter, r *http.Request) {

	category := chi.URLParam(r, "category")

	slug := chi.URLParam(r, "slug")

	blog, err := m.DB.GetBlogBySlugAndCategory(slug, category)

	if err != nil {
		log.Println(err)
		return
	}

	relatedBlogsCount := 3

	relatedBlogs, err := m.DB.GetRelatedBlogs(relatedBlogsCount, category)

	if err != nil {
		log.Println(err)
		return
	}

	data := make(map[string]interface{})

	data["blog"] = blog
	data["relatedBlogs"] = relatedBlogs

	var templateData models.TemplateData

	templateData.Data = data

	render.RenderTemplate(w, r, "blog.page.html", &templateData)
}

func (m *Repository) GetBlogsByCategory(w http.ResponseWriter, r *http.Request) {

	category := chi.URLParam(r, "category")

	blogs, err := m.DB.GetBlogsByCategory(category)

	if err != nil {
		return
	}

	data := make(map[string]interface{})

	data["blogs"] = blogs

	var templateData models.TemplateData

	templateData.Data = data

	render.RenderTemplate(w, r, "category.page.html", &templateData)
}
