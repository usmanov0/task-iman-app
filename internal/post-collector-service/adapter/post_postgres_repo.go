package adapter

import (
	"fmt"
	"github.com/jackc/pgx"
	"test-project-iman/internal/post-collector-service/domain"
)

type postRepository struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) domain.PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) Save(post *domain.Post) (id int, err error) {
	query := `INSERT INTO posts(id,user_id,title,body) 
	VALUES($1,$2,$3,$4) 
	ON CONFLICT (id) DO UPDATE 
	SET user_id = $2, title = $3, body = $4 
	RETURNING id
	`

	err = p.db.QueryRow(query, post.Id, post.UserId, post.Title, post.Body).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to save post %v", err)
	}

	return id, nil
}
