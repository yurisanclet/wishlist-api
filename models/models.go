package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID 	 			 string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string `json:"name" gorm:"not null"`
	Email      string `json:"email" gorm:"unique;not null"`
	Password   string `json:"password" gorm:"not null"`
	ProfilePic string `json:"profile_pic"`
	Wishlists []Wishlist 
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Wishlist struct {
	ID 					string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID 			string `json:"user_id"`
	Title 			string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	EventDate 	time.Time `json:"event_date"`
	Items 			[]WishItem 
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type WishItem struct {
	ID         string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	WishlistID string         `json:"wishlist_id"`
	Name       string         `json:"name" gorm:"not null"`
	Description string       `json:"description"`
	Link       string         `json:"link"`
	Price      float64        `json:"price"`
	Priority   string         `json:"priority" gorm:"not null"`
	IsBought   bool           `json:"is_bought" gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}