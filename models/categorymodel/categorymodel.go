package categorymodel

import (
	"github.com/daffalandra/go-todo-exercise/config"
	"github.com/daffalandra/go-todo-exercise/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}
