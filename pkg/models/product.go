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
	Price      int64 `json:"price" binding:"required"`
	Categories  []Category
	Color      string  `json:"color"`
	ImageUrl   string `json:"image_url" binding:"required"`
	Size       string `json:"size"`
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
	ProductID int64 `json:"product_id"`
}

func init() {
	config.Connect()
	dbp = config.GetDB()
	dbp.AutoMigrate(&Product{}, &Category{})
}

func (p *Product) CreateProduct() *Product {
	dbp.NewRecord(p)
	dbp.Create(p)
	return p
}

func (c *Category) CreateCategory() *Category {
	dbp.NewRecord(c)
	dbp.Create(c)
	return c
}

func GetProductById(Id int64) (*Product, *gorm.DB) {
	var getUser Product
	db := dbu.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}
