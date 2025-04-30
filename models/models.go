package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type Sneaker struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	BrandID   uint
	CategoryID uint
	Price     float64
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
}

type Brand struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
}

// Добавление структуры Order
type Order struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	SneakerID uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Total     float64   `gorm:"not null"`
	Status    string    `gorm:"default:'pending'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Review struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	SneakerID uint      `gorm:"not null"`
	Rating    int       `gorm:"not null"`
	Comment   string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}


type Sale struct {
	ID        uint      `gorm:"primaryKey"`
	SneakerID uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Total     float64   `gorm:"not null"`
	SaleDate  time.Time `gorm:"autoCreateTime"`
}