package models

import "time"

type Persons struct {
	ID           string     `gorm:"primarykey;type:char(36);" json:"id"`
	NamaDepan    string     `gorm:"type:varchar(45);not null" json:"nama_depan"`
	NamaBelakang string     `gorm:"type:varchar(45);not null" json:"nama_belakang"`
	NoKTP        string     `gorm:"type:varchar(16);not null" json:"no_ktp"`
	Alamat       string     `gorm:"type:varchar(255);not null" json:"alamat"`
	NoTelp       string     `gorm:"type:varchar(12);not null" json:"no_telephone"`
	Image        string     `gorm:"type:varchar(255);not null" json:"image"`
	User_id      string     `gorm:"primarykey;type:char(36);" json:"user_id"`
	CreatedAt    time.Time  `gorm:"default:null"`
	UpdatedAt    time.Time  `gorm:"default:null"`
	DeleteAt     *time.Time `gorm:"default:null"`
	Users        Users      `gorm:"foreignKey:ID;references:User_id"`
}
