package Models

import "time"

type (
	Task struct {
		ID          uint   `json:"id" gorm:"primaryKey"`
		Description string `json:"description"`
		CreatedAt   string `json:"created_at"`
		Completed   bool   `json:"completed"`
		Username    string `json:"username" gorm:"column:username;"`         // Foreign key to User
		User        User   `gorm:"foreignKey:Username;references:Username"` // Establishing the relationship
	}
	User struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		Username  string    `json:"username" gorm:"column:username;unique"`
		Password  string    `json:"password" gorm:"column:password"`
		Email     string    `json:"email" gorm:"column:email;unique"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	}
)
