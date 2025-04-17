package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type Sneaker struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string
	BrandID  uint
	CategoryID uint
	Price    float64
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
}

type Brand struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
}