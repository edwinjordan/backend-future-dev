package models

type Wilayah struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Provinsi string `gorm:"type:varchar(255)" json:"provinsi"`
	Kota     string `gorm:"type:varchar(255)" json:"kota"`
	Link     string `gorm:"type:text" json:"link"`
}
type Wilayahs []Wilayah
