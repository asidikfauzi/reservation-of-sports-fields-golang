package models

import "time"

type Bookings struct {
	ID          string     `gorm:"primarykey;type:char(36);" json:"id"`
	Person_id   string     `gorm:"type:char(36);not null" json:"person_id"`
	Lapangan_id string     `gorm:"type:char(36);not null" json:"lapangan_id"`
	Start_time  time.Time  `gorm:"not null" json:"start_time"`
	End_time    time.Time  `gorm:"not null" josn:"end_time"`
	CreatedAt   time.Time  `gorm:"default:null"`
	UpdatedAt   time.Time  `gorm:"default:null"`
	DeleteAt    *time.Time `gorm:"default:null"`
}
