package models

import "time"

type Pesans struct {
	FromUser  string    `gorm:"type:char(36);not null" json:"from_user"`
	ToUser    string    `gorm:"type:char(36);not null" json:"to_user"`
	IsiPesan  string    `gorm:"type:varchar(255);not null" json:"isi_pesan"`
	Image     string    `gorm:"type:varchar(255);not null" json:"image"`
	CreatedAt time.Time `gorm:"default:null"`
	UpdatedAt time.Time `gorm:"default:null"`
	UsersFrom Users     `gorm:"foreignKey:ID;references:FromUser"`
	UsersTo   Users     `gorm:"foreignKey:ID;references:ToUser"`
}
