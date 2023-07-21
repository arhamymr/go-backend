package database

import (
	"time"
)

type Product struct {
	ID             string    `gorm:"primaryKey;default:uuid_generate_v4()"`
	Name           string
	Description    string
	Price          float64
	ProductPicture []string `gorm:"type:text[]"`
	Discount       int      `gorm:"default:0"`
	DetailProduct  string   `gorm:"default:''"`
	Likes          int      `gorm:"default:0"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	Categories    *Category `gorm:"foreignKey:CategoryName;references:Name"`
	CategoryName  string
	Order         *Order    `gorm:"foreignKey:OrderID;references:ID"`
	OrderID       string
}

type Order struct {
	ID         string    `gorm:"primaryKey;default:uuid_generate_v4()"`
	TotalPrice int
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Products  []Product
}

type Category struct {
	Name       string    `gorm:"primaryKey"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Product   []Product
}
