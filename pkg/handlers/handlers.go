package handlers

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.page.html", &models.TemplateData{})
}

func (m *Repository) Algorithms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "algorithms.page.html", &models.TemplateData{})
}

func (m *Repository) DataStructures(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "data-structures.page.html", &models.TemplateData{})
}

func (m *Repository) BubbleSort(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "bubble-sort.page.html", &models.TemplateData{})
}

func (m *Repository) Stack(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "stack.page.html", &models.TemplateData{})
}

func (m *Repository) LinearSearch(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "linear-search.page.html", &models.TemplateData{})
}

func (m *Repository) Queue(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "queue.page.html", &models.TemplateData{})
}

func (m *Repository) InorderTraversalRecursive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "recursive-inorder-traversal.page.html", &models.TemplateData{})
}

func (m *Repository) PreOrderTraversalRecursive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "recursive-preorder-traversal.page.html", &models.TemplateData{})
}

func (m *Repository) PostOrderTraversalRecursive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "recursive-postorder-traversal.page.html", &models.TemplateData{})
}

func (m *Repository) BinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) FullBinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "full-binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) CompleteBinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "complete-binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) PerfectBinaryTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "perfect-binary-tree.page.html", &models.TemplateData{})
}

func (m *Repository) BinarySearchTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "binary-search-tree.page.html", &models.TemplateData{})
}

func (m *Repository) RangeSumBST(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "range-sum-of-bst.page.html", &models.TemplateData{})
}

func (m *Repository) RootEqualsSumofChildren(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "root-equals-sum-of-children.page.html", &models.TemplateData{})
}

func (m *Repository) MergeTwoBinaryTrees(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "merge-two-binary-trees.page.html", &models.TemplateData{})
}

func (m *Repository) SearchInBinarySearchTree(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-in-a-binary-search-tree.page.html", &models.TemplateData{})
}
