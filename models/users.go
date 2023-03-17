package models

import "time"

type Users struct {
	ID        string     `gorm:"primarykey;type:char(36);" json:"id"`
	Username  string     `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email     string     `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"password"`
	Role_id   int        `gorm:"not null" json:"role_id"`
	IsActive  bool       `gorm:"default:true"`
	CreatedAt time.Time  `gorm:"default:null"`
	UpdatedAt time.Time  `gorm:"default:null"`
	DeleteAt  *time.Time `gorm:"default:null"`
	Role      Roles      `gorm:"foreignKey:ID;references:Role_id"`
}
