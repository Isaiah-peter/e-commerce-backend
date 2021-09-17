package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	dbp *gorm.DB
)

type Product struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Desc       string `json:"description" binding:"required"`
	Price      int64  `json:"price" binding:"required"`
	Categories  []Category `json:"categories"`
	Color      string  `json:"color"`
	ImageUrl   string `json:"image_url" binding:"required"`
	Size       string `json:"size"`
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	config.Connect()
	dbp = config.GetDB()
	dbp.AutoMigrate(&Product{}, &Category{})
}
