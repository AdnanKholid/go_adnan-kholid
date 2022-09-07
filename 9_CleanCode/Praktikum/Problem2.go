package main

import "fmt"

type Kendaraan struct {
	totalRoda       int
	kecepatanPerjam int
}

type Mobil Kendaraan

func (m *Mobil) berjalan() {
	m.tambahKecepatan(10)
}

func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerjam = m.kecepatanPerjam + kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := Mobil{}
	mobilLamban.berjalan()

	fmt.Println(mobilCepat, mobilLamban)
}
