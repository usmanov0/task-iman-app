package app

import (
	"errors"
	"strconv"
	"sync"
	"test-project-iman/internal/post-collector-service/domain"
	provider "test-project-iman/internal/post-collector-service/domain"
)

const CollectPages = 50

type PostService interface {
	CollectPosts() error
}

type PostCollectorService struct {
	postRepo         domain.PostRepository
	postProviderRepo provider.PostProviderRepository
	mu               sync.Mutex
}

func NewPostService(postRepo domain.PostRepository, postProviderRepo provider.PostProviderRepository) *PostCollectorService {
	return &PostCollectorService{postRepo: postRepo,
		postProviderRepo: postProviderRepo}
}

func (uc *PostCollectorService) CollectPosts() error {
	var (
		allPosts []domain.Post
		mu       sync.Mutex
		wg       sync.WaitGroup
	)
	for page := 1; page <= CollectPages; page++ {
		wg.Add(1)

		go func(page int) {
			defer wg.Done()

			pageStr := strconv.Itoa(page)
			posts, err := uc.postProviderRepo.CollectPosts(pageStr)
			if err != nil {
				return
			}

			mu.Lock()
			allPosts = append(allPosts, posts...)
			mu.Unlock()
		}(page)
	}

	wg.Wait()
	for _, post := range allPosts {
		_, err := uc.postRepo.Save(&post)
		if err != nil {
			return errors.New("Failed to save posts to the database")
		}
	}
	return nil
}
