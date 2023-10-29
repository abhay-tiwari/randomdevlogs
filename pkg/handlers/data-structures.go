package handlers

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
)

func (m *Repository) DataStructures(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "data-structures.page.html", &models.TemplateData{})
}

func (m *Repository) Stack(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "stack.page.html", &models.TemplateData{})
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
