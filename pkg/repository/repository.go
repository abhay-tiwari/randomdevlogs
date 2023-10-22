package repository

import "github.com/abhay-tiwari/randomdevlogs/pkg/models"

type DatabaseRepo interface {
	Authenticate(email, testPassword string) (int, string, error)

	AddBlog(blog models.Blog) error
	EditBlog(blog models.Blog) error
	DeleteBlog(blogId int) error
	GetAllBlogs() ([]*models.Blog, error)
	GetBlogById(blogId int) (*models.Blog, error)
	GetBlogsByCategory(category string) ([]*models.Blog, error)
	GetBlogBySlugAndCategory(slug string, category string) (*models.Blog, error)
}
