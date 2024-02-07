package domain

type PostRepository interface {
	GetList(page, limit int) ([]Post, error)
	GetOne(postId int) (*Post, error)
	Update(postId int, title, body string) error
	Delete(postId int) error
}
