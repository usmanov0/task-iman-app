package adapter

import (
	"github.com/jackc/pgx"
	"test-project-iman/internal/collection_service/domain"
)

type postRepository struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) domain.PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) Save(post *domain.Post) (int, error) {
	query := `INSERT INTO posts(user_id,title,body) 
	VALUES($1,$2,$3) 
	RETURNING id`

	var id int
	err := p.db.QueryRow(query, post.UserId, post.Title, post.Body).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
