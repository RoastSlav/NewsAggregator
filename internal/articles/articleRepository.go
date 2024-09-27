package articles

import "NewsAggregator/database"

func insertArticle(article *Article) error {
	_, err := database.DB.NamedExec("INSERT INTO articles (title, content, source_id, source_name, author, description, url, url_to_image, published_at, created_at) VALUES (:title, :content, :source_id, :source_name, :author, :description, :url, :url_to_image, :published_at, :created_at)", article)
	return err
}

func getAllArticles() ([]Article, error) {
	var articles []Article
	err := database.DB.Select(&articles, "SELECT * FROM articles")
	return articles, err
}

func getArticleById(id int) (Article, error) {
	var article Article
	err := database.DB.Get(&article, "SELECT * FROM articles WHERE id = ?", id)
	return article, err
}
