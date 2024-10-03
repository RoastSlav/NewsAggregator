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

func getAllArticles(page int, limit int) ([]Article, error) {
	var articles []Article

	if limit == 0 {
		limit = 10
	}
	err := database.DB.Select(&articles, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id LIMIT ? OFFSET ?", limit, limit*(page-1))
	return articles, err
}

func getCategoryByName(name string) (Category, error) {
	var category Category
	err := database.DB.Get(&category, "SELECT * FROM categories WHERE name = ?", name)
	return category, err
}

func insertCategory(category *Category) error {
	_, err := database.DB.NamedExec("INSERT INTO categories (name) VALUES (:name)", category)
	return err
}

func getArticlesByTitleAndAuthor(title string, author string) ([]Article, error) {
	var article []Article
	err := database.DB.Select(&article, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id WHERE title = ? AND author = ?", title, author)
	return article, err
}

func searchArticles(article SearchArticleRequest) ([]Article, error) {
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

func getArticleById(id int) (Article, error) {
	var article Article
	err := database.DB.Get(&article, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id WHERE articles.id = ?", id)
	return article, err
}

func getArticlesByCategoryName(name string, limit int, page int) ([]Article, error) {
	var articles []Article
	err := database.DB.Select(&articles, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM articles LEFT JOIN categories ON articles.category_id = categories.id WHERE categories.name = ? LIMIT ? OFFSET ?", name, limit, limit*(page-1))
	return articles, err
}

func getCategories() ([]Category, error) {
	var categories []Category
	err := database.DB.Select(&categories, "SELECT * FROM categories")
	return categories, err
}

func addLikeToArticle(articleId int, userId int) error {
	_, err := database.DB.Exec("INSERT INTO likes (article_id, user_id) VALUES (?, ?)", articleId, userId)
	return err
}

func deleteLikeFromArticle(articleId int, userId int) error {
	_, err := database.DB.Exec("DELETE FROM likes WHERE article_id = ? AND user_id = ?", articleId, userId)
	return err
}

func checkIfUserLikedArticle(articleId int, userId int) (bool, error) {
	var count int
	err := database.DB.Get(&count, "SELECT COUNT(*) FROM likes WHERE article_id = ? AND user_id = ?", articleId, userId)
	return count > 0, err
}

func addCommentToArticle(articleId int, userId int, comment string) error {
	_, err := database.DB.Exec("INSERT INTO comments (article_id, user_id, content) VALUES (?, ?, ?)", articleId, userId, comment)
	return err
}

func isArticleInReadLater(userId int, articleId int) (bool, error) {
	var count int
	err := database.DB.Get(&count, "SELECT COUNT(*) FROM read_later WHERE user_id = ? AND article_id = ?", userId, articleId)
	return count > 0, err
}

func addArticleToReadLater(userId int, articleId int) error {
	_, err := database.DB.Exec("INSERT INTO read_later (user_id, article_id) VALUES (?, ?)", userId, articleId)
	return err
}

func removeArticleFromReadLater(userId int, articleId int) error {
	_, err := database.DB.Exec("DELETE FROM read_later WHERE user_id = ? AND article_id = ?", userId, articleId)
	return err
}

func getReadLaterArticles(userId int) ([]Article, error) {
	var articles []Article
	err := database.DB.Select(&articles, "SELECT articles.id, articles.author, articles.created_at, articles.content, articles.description, articles.source_id ,articles.source_name, articles.title, articles.published_at, articles.url, articles.url_to_image, categories.name AS category FROM read_later LEFT JOIN articles ON read_later.article_id = articles.id LEFT JOIN categories ON articles.category_id = categories.id WHERE read_later.user_id = ?", userId)
	return articles, err
}

func getCommentsForArticle(articleId int) ([]Comment, error) {
	var comments []Comment
	err := database.DB.Select(&comments, "SELECT comments.id, comments.article_id, comments.user_id, comments.content, comments.created_at FROM comments WHERE comments.article_id = ?", articleId)
	return comments, err
}
