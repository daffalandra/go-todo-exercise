package categorymodel

import (
	"time"

	"github.com/daffalandra/go-todo-exercise/entities"
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) ([]entities.Category, error) {
	var categories []entities.Category
	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func Create(db *gorm.DB, name string) error {
	category := entities.Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return db.Create(&category).Error
}

func GetByID(db *gorm.DB, id uint) (*entities.Category, error) {
	var category entities.Category
	err := db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func Update(db *gorm.DB, id uint, name string) error {
	category := entities.Category{
		ID:        id,
		Name:      name,
		UpdatedAt: time.Now(),
	}
	return db.Save(&category).Error
}

func Delete(db *gorm.DB, id uint) error {
	return db.Delete(&entities.Category{}, id).Error
}
