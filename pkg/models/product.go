package models

type Product struct {
	Name       string `json:"name" binding:"required"`
	Desc       string `json:"description" binding:"required"`
	Price      int64  `json:"price" binding:"required"`
	CountStock int64  `json:"count_stock" binding:"required"`
	ImageUrl   string `json:"image_url" binding:"required"`
}
