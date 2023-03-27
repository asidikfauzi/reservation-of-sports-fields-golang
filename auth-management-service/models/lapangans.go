package models

import "time"

type Lapangans struct {
	ID               string        `gorm:"primarykey;type:char(36);" json:"id"`
	Nama             string        `gorm:"type:varchar(45);not null" json:"nama"`
	Alamat           string        `gorm:"type:varchar(255);not null" json:"alamat"`
	Harga_per_jam    string        `gorm:"type:varchar(45);not null" json:"harga_per_jam"`
	Tipe_lapangan_id string        `gorm:"primarykey;type:char(36);" json:"tipe_lapangan_id"`
	Person_id        string        `gorm:"primarykey;type:char(36);" json:"person_id"`
	CreatedAt        time.Time     `gorm:"default:null"`
	UpdatedAt        time.Time     `gorm:"default:null"`
	DeleteAt         *time.Time    `gorm:"default:null"`
	TipeLapangan     TipeLapangans `gorm:"foreignKey:ID;references:Tipe_lapangan_id"`
	Person           Persons       `gorm:"foreignKey:ID;references:Person_id"`
}
