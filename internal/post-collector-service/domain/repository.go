package domain

type PostRepository interface {
	Save(post *Post) (int, error)
}

type PostProviderRepository interface {
	FetchPosts(page string) ([]Post, error)
}
