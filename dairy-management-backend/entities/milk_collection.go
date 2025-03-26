// entities/milk_collection.go - Milk collection entity
package entities

import "time"

type MilkCollection struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CowID     uint      `json:"cow_id"`
	Quantity  float64   `json:"quantity"` // Liters of milk collected
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
