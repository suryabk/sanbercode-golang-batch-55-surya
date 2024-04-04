package main

import (
	"fmt"
	"strings"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type displayData interface {
	displayPhone() string
}

func main() {

	// Soal 1
	segitiga := segitigaSamaSisi{alas: 5, tinggi: 4}
	persegi := persegiPanjang{panjang: 6, lebar: 4}
	tabung := tabung{jariJari: 3, tinggi: 5}
	balok := balok{panjang: 3, lebar: 4, tinggi: 5}

	fmt.Println("Luas segitiga:", segitiga.luas())
	fmt.Println("Keliling segitiga:", segitiga.keliling())
	fmt.Println("Luas persegi panjang:", persegi.luas())
	fmt.Println("Keliling persegi panjang:", persegi.keliling())
	fmt.Printf("Volume tabung: %.2f\n", tabung.volume())
	fmt.Println("Luas permukaan tabung:", tabung.luasPermukaan())
	fmt.Println("Volume balok:", balok.volume())
	fmt.Println("Luas permukaan balok:", balok.luasPermukaan())
	fmt.Printf("\n\n")

	// Soal 2
	smartphone := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(smartphone.display())
	fmt.Printf("\n\n")

	// Soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
	fmt.Printf("\n\n")

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefixString := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := sumNumbers(angkaPertama...)
	total += sumNumbers(angkaKedua...)

	fmt.Printf("%s%d + %d + %d + %d = %d\n", prefixString, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
}

func sumNumbers(angka ...int) int {
	total := 0
	for _, num := range angka {
		total += num
	}
	return total
}

func luasPersegi(sisi int, description bool) interface{} {
	if sisi == 0 {
		if description {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}

	luas := sisi * sisi

	if description {
		return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	} else {
		return luas
	}
}

func (p phone) display() string {
	return fmt.Sprintf("name: %s\nbrand: %s\nyear: %d\ncolors: %s", p.name, p.brand, p.year, strings.Join(p.colors, ", "))
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return 3.14 * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return (2 * 3.14 * t.jariJari * t.tinggi) + (2 * 3.14 * t.jariJari * t.jariJari)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2*float64(b.panjang*b.lebar) + 2*float64(b.panjang*b.tinggi) + 2*float64(b.lebar*b.tinggi)
}
