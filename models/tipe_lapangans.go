package models

import "time"

type TipeLapangans struct {
	ID        string     `gorm:"primarykey;type:char(36);" json:"id"`
	Nama      string     `gorm:"type:varchar(45);not null" json:"nama"`
	CreatedAt time.Time  `gorm:"default:null"`
	UpdatedAt time.Time  `gorm:"default:null"`
	DeleteAt  *time.Time `gorm:"default:null"`
}
