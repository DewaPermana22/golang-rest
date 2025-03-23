package models

type Category struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"unique;not null;type:varchar(255)" json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products"`
}
