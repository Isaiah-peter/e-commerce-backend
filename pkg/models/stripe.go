package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)
type Charge struct {
	gorm.Model
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Charge{})
}

func (c *Charge) CreateStripe() *Charge {
	db.NewRecord(c)
	db.Create(c)
	return c
}
