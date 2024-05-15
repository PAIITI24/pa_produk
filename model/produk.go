package model

import (
	"time"
)

type Produk struct {
	Id             int              `json:"id" gorm:"primaryKey"`
	NamaProduk     string           `json:"nama_produk" gorm:"type:varchar(50);not null"`
	Harga          float64          `json:"harga" gorm:"type:float;not null"`
	Gambar         string           `json:"gambar" gorm:"type:text"`
	Deskripsi      string           `json:"deskripsi" gorm:"type:text;not null"`
	CreatedAt      time.Time        `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	KategoriProduk []KategoriProduk `json:"kategori_produk,omitempty" gorm:"many2many:kategorisasi"`
}

type KategoriProduk struct {
	Id           int       `json:"id" gorm:"PrimaryKey"`
	NamaKategori string    `json:"nama_kategori_produk" gorm:"type:varchar(50);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Produk       []Produk  `json:"produk,omitempty" gorm:"many2many:kategorisasi"`
}
