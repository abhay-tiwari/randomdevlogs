package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *postgresDBRepo) AuthenticateUser(email string, password string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int

	var hashedPassword string

	query := `select email, password from users where email = $1`

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(&id, &hashedPassword)

	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("Incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

func (m *postgresDBRepo) AddBlog(blog models.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	statement := `insert into blogs(title, meta_description, og_title, og_description, slug, content, created_by, category, tags, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := m.DB.ExecContext(ctx, statement, blog.Title, blog.MetaDescription, blog.OgTitle, blog.OgDescription, blog.Slug, blog.Content, blog.CreatedBy, blog.Category, blog.Tags, time.Now(), time.Now())

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *postgresDBRepo) GetAllBlogs() ([]*models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	blogs := make([]*models.Blog, 0)

	query := `select title, meta_description, og_title, og_description, slug, content, created_by, category, created_at, updated_at, tags from blogs`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return blogs, err
	}

	defer rows.Close()

	for rows.Next() {
		i := new(models.Blog)

		err := rows.Scan(
			&i.Title,
			&i.MetaDescription,
			&i.OgTitle,
			&i.OgDescription,
			&i.Slug,
			&i.Content,
			&i.CreatedBy,
			&i.Category,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Tags,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		blogs = append(blogs, i)
	}

	return blogs, nil
}

func (m *postgresDBRepo) GetBlogBySlugAndCategory(slug string, category string) (*models.Blog, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	blog := new(models.Blog)

	row := m.DB.QueryRowContext(ctx, "SELECT id, title, meta_description, og_title, og_description, slug, content, created_by, category, tags FROM blogs WHERE slug = $1", slug)

	err := row.Scan(
		&blog.ID,
		&blog.Title,
		&blog.MetaDescription,
		&blog.OgTitle,
		&blog.OgDescription,
		&blog.Slug,
		&blog.Content,
		&blog.CreatedBy,
		&blog.Category,
		&blog.Tags,
	)

	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (m *postgresDBRepo) GetBlogsByCategory(category string) ([]*models.Blog, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	blogs := make([]*models.Blog, 0)

	query := `select title, meta_description, og_title, og_description, slug, content, created_by, category, created_at, updated_at, tags from blogs where category=$1`

	rows, err := m.DB.QueryContext(ctx, query, category)

	defer rows.Close()

	if err != nil {
		return blogs, err
	}

	for rows.Next() {
		i := new(models.Blog)

		err := rows.Scan(
			&i.Title,
			&i.MetaDescription,
			&i.OgTitle,
			&i.OgDescription,
			&i.Slug,
			&i.Content,
			&i.CreatedBy,
			&i.Category,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Tags,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		blogs = append(blogs, i)
	}

	return blogs, nil
}
