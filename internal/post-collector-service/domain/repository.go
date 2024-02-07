package domain

type PostRepository interface {
	Save(post *Post) (int, error)
}

type PostProviderRepository interface {
	CollectPosts(page string) ([]Post, error)
}
