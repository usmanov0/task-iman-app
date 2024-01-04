package domain

type PostRepository interface {
	Save(post *Post) (int, error)
}
