package utils

import (
	"fmt"
	"strings"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	volume() float64
	LuasPermukaan() float64
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type DisplayData interface {
	DisplayPhone() string
}

func SumNumbers(angka ...int) int {
	total := 0
	for _, num := range angka {
		total += num
	}
	return total
}

func LuasPersegi(sisi int, description bool) interface{} {
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

func (p Phone) Display() string {
	return fmt.Sprintf("name: %s\nbrand: %s\nyear: %d\ncolors: %s", p.Name, p.Brand, p.Year, strings.Join(p.Colors, ", "))
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return 3.14 * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return (2 * 3.14 * t.JariJari * t.Tinggi) + (2 * 3.14 * t.JariJari * t.JariJari)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2*float64(b.Panjang*b.Lebar) + 2*float64(b.Panjang*b.Tinggi) + 2*float64(b.Lebar*b.Tinggi)
}
