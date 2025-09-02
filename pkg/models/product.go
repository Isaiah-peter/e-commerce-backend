package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type Product struct {
	gorm.Model
	Title      string `json:"title" binding:"required"`
	Desc       string `json:"description" binding:"required"`
	Price      int64  `json:"price" binding:"required"`
	Categories []Category
	Color      []Color
	ImageUrl   string `json:"image_url" binding:"required"`
	Size       []Size
	InStock    int64 ` json:"in_stock"`
}

type Category struct {
	gorm.Model
	Name      string `json:"name"`
	ProductID int64  `json:"product_id"`
}

type Color struct {
	gorm.Model
	Name      string `json:"name"`
	ProductID int64  `json:"product_id"`
}

type Size struct {
	gorm.Model
	Name      string `json:"name"`
	ProductID int64  `json:"product_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{}, &Category{}, &Size{}, &Color{})
}

func (p *Product) CreateProduct() *Product {
	db.NewRecord(p)
	db.Create(p)
	return p
}

func (c *Category) CreateCategory() *Category {
	db.NewRecord(c)
	db.Create(c)
	return c
}

func GetProductById(Id int64) (*Product, *gorm.DB) {
	var GetProduct Product
	db := db.Where("ID=?", Id).Preload("Color").Preload("Size").Preload("Categories").Find(&GetProduct)
	return &GetProduct, db
}
