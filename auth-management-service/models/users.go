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

type Users_register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role_id  int    `json:"role_id" binding:"required"`
}

type Users_login struct {
	ID      int    `gorm:"primarykey;" json:"id"`
	Email   string `json:"email" binding:"required,email"`
	Role_id int    `json:"role_id" binding:"required"`
	Token   string `json:"token"`
}

type Users_update struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Customers_response struct {
	Code    string
	Message string
	Status  string
}
