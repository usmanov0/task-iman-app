package domain

type Post struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
