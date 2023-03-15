package models

import "time"

type Ratings struct {
	ID          string    `gorm:"primarykey;type:char(36);" json:"id"`
	Person_id   string    `gorm:"type:char(36);not null" json:"person_id"`
	Lapangan_id string    `gorm:"type:char(36);not null" json:"lapangan_id"`
	Jumlah      int       `gorm:"not null" json:"jumlah"`
	Ulasan      string    `gorm:"type:varchar(255);not null" json:"ulasan"`
	CreatedAt   time.Time `gorm:"default:null"`
	UpdatedAt   time.Time `gorm:"default:null"`
}
