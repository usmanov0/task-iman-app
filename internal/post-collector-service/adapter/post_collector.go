package adapter

import (
	"encoding/json"
	"github.com/jackc/pgx"
	"net/http"
	"test-project-iman/internal/post-collector-service/domain"
)

type postProviderRepo struct {
	db *pgx.Conn
}

func NewPostCollectorRepository(db *pgx.Conn) domain.PostProviderRepository {
	return &postProviderRepo{db: db}
}

type Pagination struct {
	Total int   `json:"total"`
	Pages int   `json:"pages"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Links Links `json:"links"`
}

type Links struct {
	Previous *string `json:"previous"`
	Current  string  `json:"current"`
	Next     string  `json:"next"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type ApiResponse struct {
	Meta Meta          `json:"meta"`
	Data []domain.Post `json:"data"`
}

func (p *postProviderRepo) CollectPosts(page string) ([]domain.Post, error) {
	res, err := http.Get("https://gorest.co.in/public/v1/posts?page=" + page)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response ApiResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	var postList []domain.Post
	for _, post := range response.Data {
		postList = append(postList, domain.Post{
			Id:     post.Id,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		})
	}
	return postList, nil
}
