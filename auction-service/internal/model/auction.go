package model

import (
	"gorm.io/gorm"
)

type Auction struct {
	gorm.Model
	Item   string `json:"item"`
	UserID uint   `json:"user_id"`
}
