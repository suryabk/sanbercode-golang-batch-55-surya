package main

import (
	"fmt"
	"strings"
)

type Buah struct {
	nama, warna string
	adaBiji     bool
	harga       int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title, genre   string
	duration, year int
}

func main() {

	// Soal 1
	fruits := []Buah{
		{nama: "Nanas", warna: "Kuning", adaBiji: false, harga: 9000},
		{nama: "Jeruk", warna: "Oranye", adaBiji: true, harga: 8000},
		{nama: "Semangka", warna: "Hijau & Merah", adaBiji: true, harga: 10000},
		{nama: "Pisang", warna: "Kuning", adaBiji: false, harga: 5000},
	}

	for index, fruit := range fruits {
		isAdaBiji := "Tidak"
		if fruit.adaBiji {
			isAdaBiji = "Ada"
		}

		fmt.Printf("%d. Nama Buah : %s\n   Warna : %s\n   Ada Biji : %s\n   Harga : %d\n\n", index+1, fruit.nama, fruit.warna, isAdaBiji, fruit.harga)
	}
	fmt.Printf("\n\n")

	// Soal 2
	segitiga := segitiga{alas: 2, tinggi: 3}
	persegi := persegi{sisi: 4}
	persegiPanjang := persegiPanjang{panjang: 4, lebar: 6}
	fmt.Println("Luas Segitiga : ", luasSegitiga(segitiga.alas, segitiga.tinggi))
	fmt.Println("Luas Persegi : ", luasPersegi(persegi))
	fmt.Println("Luas Persegi Panjang : ", luasPersegiPanjang(persegiPanjang))
	fmt.Printf("\n\n")

	// Soal 3
	var smartphone = phone{name: "Hp Murah", colors: []string{"putih"}, brand: "Samsung", year: 2023}

	tambahWarna(&smartphone, "tosca")
	fmt.Println("Warna Smartphone : ", strings.Join(smartphone.colors, ", "))
	fmt.Printf("\n\n")

	// Soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for index, film := range dataFilm {
		fmt.Printf("%d. title : %s\n   duration : %d jam\n   genre : %s\n   year : %d\n\n", index+1, film.title, film.duration, film.genre, film.year)
	}

}

func tambahWarna(phone *phone, warna string) {
	*&phone.colors = append(*&phone.colors, warna)
}

func luasSegitiga(alas, tinggi int) int {
	return (alas * tinggi) / 2
}

func luasPersegi(p persegi) int {
	return p.sisi * p.sisi

}

func luasPersegiPanjang(p persegiPanjang) int {
	return p.panjang * p.lebar
}

func tambahDataFilm(name string, duration int, genre string, year int, film *[]movie) {
	*film = append(*film, movie{title: name, duration: duration, genre: genre, year: year})
}
