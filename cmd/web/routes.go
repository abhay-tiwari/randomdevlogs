package main

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Index)
	mux.Get("/all-blogs", handlers.Repo.GetAllBlogs)

	mux.Get("/add-blog", handlers.Repo.AddBlog)
	mux.Post("/add-blog", handlers.Repo.SubmitBlog)

	mux.Get("/login", handlers.Repo.GetLoginPage)
	mux.Post("/login", handlers.Repo.Login)
	mux.Post("/admin", handlers.Repo.Admin)
	/*	mux.Route("/algorithms", func(r chi.Router) {
			r.Get("/bubble-sort", handlers.Repo.BubbleSort)
			r.Get("/linear-search", handlers.Repo.LinearSearch)
			r.Get("/", handlers.Repo.Algorithms)
		})
	*/
	mux.Route("/data-structures", func(r chi.Router) {
		r.Get("/stack", handlers.Repo.Stack)
		r.Get("/queue", handlers.Repo.Queue)
		r.Get("/recursive-inorder-traversal", handlers.Repo.InorderTraversalRecursive)
		r.Get("/recursive-preorder-traversal", handlers.Repo.PreOrderTraversalRecursive)
		r.Get("/recursive-postorder-traversal", handlers.Repo.PostOrderTraversalRecursive)
		r.Get("/binary-tree", handlers.Repo.BinaryTree)
		r.Get("/full-binary-tree", handlers.Repo.FullBinaryTree)
		r.Get("/complete-binary-tree", handlers.Repo.CompleteBinaryTree)
		r.Get("/perfect-binary-tree", handlers.Repo.PerfectBinaryTree)
		r.Get("/binary-search-tree", handlers.Repo.BinarySearchTree)
		r.Get("/", handlers.Repo.DataStructures)
	})

	mux.Get("/{category}/{slug}", handlers.Repo.GetBlogBySlugAndCategory)

	mux.Get("/{category}", handlers.Repo.GetBlogsByCategory)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
