package adapter

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"test-project-iman/internal/post-service/domain"
)

type postRepository struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) domain.PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) GetList(page, limit int) ([]domain.Post, error) {
	query := `SELECT p.id, p.user_id, p.title, p.body FROM posts p WHERE p.page = $1  LIMIT $2`

	rows, err := p.db.Query(query, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get all posts %v", err)
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %v", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *postRepository) GetOne(postId int) (*domain.Post, error) {
	var post domain.Post
	query := `SELECT p.id, p.user_id, p.title, p.body FROM posts p WHERE p.id = $1`

	err := p.db.QueryRow(query, postId).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to execute query %v", err)
	}
	return &post, nil
}

func (p *postRepository) Update(postId int, title, body string) error {
	query := `UPDATE posts SET title = $2, body = $3 WHERE id = $1`

	_, err := p.db.Query(query, postId, title, body)
	if err != nil {
		fmt.Errorf("failed to update post %v", err)
	}
	return nil
}

func (p *postRepository) Delete(postId int) error {
	query := `DELETE FROM posts WHERE id = $1`

	result, err := p.db.Exec(query, postId)
	if err != nil {
		return fmt.Errorf("failed to delete post %v", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}
