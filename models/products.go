package models

type Product struct {
	ID         uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string   `gorm:"not null" json:"name"`
	Price      float64  `json:"price"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}

