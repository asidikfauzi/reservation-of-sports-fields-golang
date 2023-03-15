package models

import "time"

type Users struct {
	ID        string     `gorm:"primarykey;type:char(36);" json:"id"`
	Username  string     `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email     string     `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"password"`
	Role      string     `gorm:"type:varchar(10);not null" json:"role"`
	IsActive  bool       `gorm:"default:true"`
	CreatedAt time.Time  `gorm:"default:null"`
	UpdatedAt time.Time  `gorm:"default:null"`
	DeleteAt  *time.Time `gorm:"default:null"`
}
