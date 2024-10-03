package articles

import (
	"NewsAggregator/internal/database"
	"strconv"
)

func insertArticle(article *Article, categoryId int) error {
	_, err := database.DB.NamedExec("INSERT IGNORE INTO articles (title, content, source_id, source_name, author, description, url, url_to_image, published_at, category_id) VALUES (:title, :content, :source_id, :source_name, :author, :description, :url, :url_to_image, :published_at, :categoryId)",
		map[string]interface{}{
			"title":        article.Title,
			"content":      article.Content,
			"source_id":    article.SourceID,
			"source_name":  article.SourceName,
			"author":       article.Author,
			"description":  article.Description,
			"url":          article.URL,
			"url_to_image": article.URLToImage,
			"published_at": article.PublishedAt,
			"categoryId":   categoryId,
		})
	return err
}

func GetAllArticles(page int, limit int) ([]Article, error) {
	var articles []Article

	if limit == 0 {
		limit = 10
	}
	err := database.DB.Select(&articles, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id LIMIT ? OFFSET ?", limit, limit*(page-1))
	return articles, err
}

func GetCategoryByName(name string) (Category, error) {
	var category Category
	err := database.DB.Get(&category, "SELECT * FROM categories WHERE name = ?", name)
	return category, err
}

func InsertCategory(category *Category) error {
	_, err := database.DB.NamedExec("INSERT INTO categories (name) VALUES (:name)", category)
	return err
}

func GetArticlesByTitleAndAuthor(title string, author string) ([]Article, error) {
	var article []Article
	err := database.DB.Select(&article, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id WHERE title = ? AND author = ?", title, author)
	return article, err
}

func SearchArticles(article SearchArticleRequest) ([]Article, error) {
	querry := "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id WHERE "

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

	if !article.PublishedFrom.IsZero() {
		if requiresAnd {
			querry += " AND "
		}
		querry += "published_at >= '" + article.PublishedFrom.String() + "'"
		requiresAnd = true
	}

	if !article.PublishedTo.IsZero() {
		if requiresAnd {
			querry += " AND "
		}
		querry += "published_at <= '" + article.PublishedTo.String() + "'"
	}

	if article.Limit == 0 {
		article.Limit = 10
	}
	querry += " LIMIT " + strconv.Itoa(article.Limit) + " OFFSET " + strconv.Itoa(article.Limit*(article.Page-1))

	var articles []Article
	err := database.DB.Select(&articles, querry)
	return articles, err
}

func GetArticleById(id int) (Article, error) {
	var article Article
	err := database.DB.Get(&article, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id WHERE articles.id = ?", id)
	return article, err
}

func GetArticlesByCategoryName(name string, limit int, page int) ([]Article, error) {
	var articles []Article
	err := database.DB.Select(&articles, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id WHERE categories.name = ? LIMIT ? OFFSET ?", name, limit, limit*(page-1))
	return articles, err
}
