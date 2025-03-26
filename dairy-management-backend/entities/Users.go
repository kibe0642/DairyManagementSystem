package entities

import "time"

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	IsAdmin   bool      `gorm:"type:bit;default:0" json:"is_admin"`
}
