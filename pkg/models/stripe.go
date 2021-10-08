package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	dbss *gorm.DB
)

type Charge struct {
	gorm.Model
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func init() {
	config.Connect()
	dbss = config.GetDB()
	dbss.AutoMigrate(&Charge{})
}

func (c *Charge) CreateStripe() *Charge {
	dbss.NewRecord(c)
	dbss.Create(c)
	return c
}
