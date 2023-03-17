package models

import "time"

type Roles struct {
	ID        int        `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name      string     `gorm:"type:varchar(10);not null" json:"name"`
	CreatedAt time.Time  `gorm:"default:null"`
	UpdatedAt time.Time  `gorm:"default:null"`
	DeleteAt  *time.Time `gorm:"default:null"`
}
