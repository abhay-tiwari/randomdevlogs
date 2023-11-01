package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
)

func (m *postgresDBRepo) AddBlog(blog models.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	statement := `insert into blogs(title, meta_description, og_title, og_description, slug, content, created_by, category, tags, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := m.DB.ExecContext(ctx, statement, blog.Title, blog.MetaDescription, blog.OgTitle, blog.OgDescription, blog.Slug, blog.Content, blog.CreatedBy, blog.Category, blog.Tags, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) EditBlog(blog models.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `UPDATE blogs SET title=$1, meta_description=$2, og_title=$3, og_description=$4, slug=$5, content=$6, category=$7, tags=$8, created_at=$9, created_by=$10, updated_at=$11 where slug=$12`

	_, err := m.DB.ExecContext(ctx, statement, blog.Title, blog.MetaDescription, blog.OgTitle, blog.OgDescription, blog.Slug, blog.Content, blog.Category, blog.Tags, blog.CreatedAt, blog.CreatedBy, time.Now(), blog.Slug)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) DeleteBlog(blogId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `DELETE from blogs WHERE id=$1`

	_, err := m.DB.ExecContext(ctx, statement, blogId)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) GetAllBlogs() ([]*models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	blogs := make([]*models.Blog, 0)

	query := `select id, title, meta_description, og_title, og_description, slug, content, created_by, category, created_at, updated_at, tags from blogs`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return blogs, err
	}

	defer rows.Close()

	for rows.Next() {
		i := new(models.Blog)

		err := rows.Scan(
			&i.ID,
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

func (m *postgresDBRepo) GetBlogById(blogId int) (*models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	blog := new(models.Blog)

	row := m.DB.QueryRowContext(ctx, "SELECT id, title, meta_description, og_title, og_description, slug, content, created_by, category, tags FROM blogs WHERE id = $1", blogId)

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

func (m *postgresDBRepo) GetRelatedBlogs(count int, category string) ([]*models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, meta_description, og_title, og_description, slug, content, created_by, category, tags FROM blogs WHERE category = $1 ORDER BY created_at desc limit $2`

	blogs := make([]*models.Blog, 0)

	rows, err := m.DB.QueryContext(ctx, query, category, count)

	if err != nil {
		return blogs, err
	}

	for rows.Next() {
		i := new(models.Blog)

		err := rows.Scan(&i.ID, &i.Title, &i.MetaDescription, &i.OgTitle, &i.OgDescription, &i.Slug, &i.Content, &i.CreatedBy, &i.Category, &i.Tags)

		if err != nil {
			return blogs, err
		}

		blogs = append(blogs, i)
	}

	return blogs, nil
}
