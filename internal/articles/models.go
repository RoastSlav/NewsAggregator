package articles

import "time"

type Article struct {
	ID          int       `db:"id"`
	SourceID    string    `json:"source.id" db:"source_id"`
	SourceName  string    `json:"source.name" db:"source_name"`
	Author      string    `json:"author" db:"author"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	URL         string    `json:"url" db:"url"`
	URLToImage  string    `json:"urlToImage" db:"url_to_image"`
	PublishedAt time.Time `json:"publishedAt" db:"published_at"`
	Content     string    `json:"content" db:"content"`
	CreatedAt   time.Time `db:"created_at"`
}

type SearchArticle struct {
	SourceID      string    `json:"source.id"`
	SourceName    string    `json:"source.name"`
	Author        string    `json:"author"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	PublishedFrom time.Time `json:"publishedFrom"`
	PublishedTo   time.Time `json:"publishedTo"`
}

type NewsAPIResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
