package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/driver"
	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
	"github.com/abhay-tiwari/randomdevlogs/pkg/repository"
	"github.com/abhay-tiwari/randomdevlogs/pkg/repository/dbrepo"
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
	render.RenderTemplate(w, r, "index.page.html", &models.TemplateData{
		Active: "index",
	})
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

func (m *Repository) GetSitemap(w http.ResponseWriter, r *http.Request) {
	fileContent, err := ioutil.ReadFile("sitemap.xml")

	if err != nil {
		log.Println("Could not read sitemap")
		return
	}

	w.Write(fileContent)
	return
}
