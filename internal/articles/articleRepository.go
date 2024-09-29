package articles

import "NewsAggregator/database"

func insertArticle(article *Article) error {
	_, err := database.DB.NamedExec("INSERT IGNORE INTO articles (title, content, source_id, source_name, author, description, url, url_to_image, published_at, created_at) VALUES (:title, :content, :source_id, :source_name, :author, :description, :url, :url_to_image, :published_at, :created_at)", article)
	return err
}

func GetAllArticles() ([]Article, error) {
	var articles []Article
	err := database.DB.Select(&articles, "SELECT * FROM articles")
	return articles, err
}

func SearchArticles(article SearchArticle) ([]Article, error) {
	querry := "SELECT * FROM articles WHERE "

	requiresAnd := false

	if article.SourceID != "" {
		querry += "source_id LIKE '%" + article.SourceID + "%'"
		requiresAnd = true
	}

	if article.SourceName != "" {
		if requiresAnd {
			querry += " AND "
		}
		querry += "source_name LIKE '%" + article.SourceName + "%'"
		requiresAnd = true
	}

	if article.Author != "" {
		if requiresAnd {
			querry += " AND "
		}
		querry += "author LIKE '%" + article.Author + "%'"
		requiresAnd = true
	}

	if article.Title != "" {
		if requiresAnd {
			querry += " AND "
		}
		querry += "title LIKE '%" + article.Title + "%'"
		requiresAnd = true
	}

	if article.Description != "" {
		if requiresAnd {
			querry += " AND "
		}
		querry += "description LIKE '%" + article.Description + "%'"
		requiresAnd = true
	}

	if article.PublishedAt != "" {
		if requiresAnd {
			querry += " AND "
		}
		querry += "published_at = '" + article.PublishedAt + "'"
	}

	var articles []Article
	err := database.DB.Select(&articles, querry)
	return articles, err
}

func GetArticleById(id int) (Article, error) {
	var article Article
	err := database.DB.Get(&article, "SELECT * FROM articles WHERE id = ?", id)
	return article, err
}
