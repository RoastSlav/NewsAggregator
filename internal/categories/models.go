package categories

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type PagedRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
