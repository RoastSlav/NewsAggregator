package categories

import "NewsAggregator/internal/database"

func GetCategories() ([]Category, error) {
	var categories []Category
	err := database.DB.Select(&categories, "SELECT * FROM categories")
	return categories, err
}
