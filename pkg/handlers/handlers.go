package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/driver"
	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
	"github.com/abhay-tiwari/randomdevlogs/pkg/repository"
	"github.com/abhay-tiwari/randomdevlogs/pkg/repository/dbrepo"
	"github.com/go-chi/chi/v5"
	"github.com/gomarkdown/markdown"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostGresRepo(db.SQL, a),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "index.page.html", &models.TemplateData{})
}

func (m *Repository) Algorithms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "algorithms.page.html", &models.TemplateData{})
}

func (m *Repository) DataStructures(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "data-structures.page.html", &models.TemplateData{})
}

func (m *Repository) BubbleSort(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "bubble-sort.page.html", &models.TemplateData{})
}

func (m *Repository) Stack(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "stack.page.html", &models.TemplateData{})
}

func (m *Repository) LinearSearch(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "linear-search.page.html", &models.TemplateData{})
}

func (m *Repository) Queue(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "queue.page.html", &models.TemplateData{})
}

func (m *Repository) InorderTraversalRecursive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "recursive-inorder-traversal.page.html", &models.TemplateData{})
}

func (m *Repository) PreOrderTraversalRecursive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "recursive-preorder-traversal.page.html", &models.TemplateData{})
}

func (m *Repository) PostOrderTraversalRecursive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "recursive-postorder-traversal.page.html", &models.TemplateData{})
}

func (m *Repository) BinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) FullBinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "full-binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) CompleteBinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "complete-binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) PerfectBinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "perfect-binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) BinarySearchTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "binary-search-tree.page.html", &models.TemplateData{})
}

func (m *Repository) RangeSumBST(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "range-sum-of-bst.page.html", &models.TemplateData{})
}

func (m *Repository) RootEqualsSumofChildren(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "root-equals-sum-of-children.page.html", &models.TemplateData{})
}

func (m *Repository) MergeTwoBinaryTrees(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "merge-two-binary-trees.page.html", &models.TemplateData{})
}

func (m *Repository) SearchInBinarySearchTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-in-a-binary-search-tree.page.html", &models.TemplateData{})
}

func (m *Repository) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, _ := m.DB.GetAllBlogs()

	data := make(map[string]interface{})

	data["blogs"] = blogs

	var templateData models.TemplateData

	templateData.Data = data

	render.RenderTemplate(w, r, "all-blogs.page.html", &templateData)
}

func (m *Repository) AddBlog(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "add-blog.page.html", &models.TemplateData{})
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
		log.Println("file upload failed")
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

	m.DB.AddBlog(blog)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func (m *Repository) GetBlogBySlugAndCategory(w http.ResponseWriter, r *http.Request) {

	category := chi.URLParam(r, "category")

	slug := chi.URLParam(r, "slug")

	blog, err := m.DB.GetBlogBySlugAndCategory(slug, category)

	if err != nil {
		log.Println(err)

		return
	}

	data := make(map[string]interface{})

	data["blog"] = blog

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
