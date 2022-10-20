package models

import "time"

type Item struct {
	ID          uint
	Nama        string
	Description string
	Stok        uint
	Harga       float32
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemInput struct {
	Nama        string
	Description string
	Stok        uint
	Harga       float32
	CategoryID  uint
}
