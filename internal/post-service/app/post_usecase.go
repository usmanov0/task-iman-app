package app

import (
	"fmt"
	"test-project-iman/internal/post-service/domain"
)

type PostService interface {
	GetAllPosts(page, limit int) ([]domain.Post, error)
	GetOne(postId int) (*domain.Post, error)
	Update(postId int, title, body string) error
	Delete(postId int) error
}

func NewPostUseCase(postRepo domain.PostRepository) PostService {
	return &postUseCase{postRepo: postRepo}
}

type postUseCase struct {
	postRepo domain.PostRepository
}

func (p *postUseCase) GetAllPosts(page, limit int) ([]domain.Post, error) {
	posts, err := p.postRepo.GetList(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get all posts: %v", err)
	}
	return posts, nil
}

func (p *postUseCase) GetOne(postId int) (*domain.Post, error) {
	post, err := p.postRepo.GetOne(postId)
	if err != nil {
		return nil, fmt.Errorf("failed to get post %v", err)
	}
	return post, nil
}

func (p *postUseCase) Update(postId int, title, body string) error {
	err := p.postRepo.Update(postId, title, body)
	if err != nil {
		return fmt.Errorf("failed to update posts %v", err)
	}

	return nil
}

func (p *postUseCase) Delete(postId int) error {
	err := p.postRepo.Delete(postId)
	if err != nil {
		return fmt.Errorf("can't delete post %v", err)
	}
	return nil
}
