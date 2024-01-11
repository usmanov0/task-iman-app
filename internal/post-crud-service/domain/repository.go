package domain

type PostCrudRepository interface {
	GetList(page, limit int) ([]Post, error)
	GetOne(postId int) (*Post, error)
	Update(postId int, title, body string) (*PostUpdateResponse, error)
	Delete(postId int) error
}
