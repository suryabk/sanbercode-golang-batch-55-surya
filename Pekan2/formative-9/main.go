package main

import (
	"fmt"
	"formative-9/utils"
)

func main() {

	// Soal 1
	segitiga := utils.SegitigaSamaSisi{Alas: 5, Tinggi: 4}
	persegi := utils.PersegiPanjang{Panjang: 6, Lebar: 4}
	tabung := utils.Tabung{JariJari: 3, Tinggi: 5}
	balok := utils.Balok{Panjang: 3, Lebar: 4, Tinggi: 5}

	fmt.Println("Luas segitiga:", segitiga.Luas())
	fmt.Println("Keliling segitiga:", segitiga.Keliling())
	fmt.Println("Luas persegi panjang:", persegi.Luas())
	fmt.Println("Keliling persegi panjang:", persegi.Keliling())
	fmt.Printf("Volume tabung: %.2f\n", tabung.Volume())
	fmt.Println("Luas permukaan tabung:", tabung.LuasPermukaan())
	fmt.Println("Volume balok:", balok.Volume())
	fmt.Println("Luas permukaan balok:", balok.LuasPermukaan())
	fmt.Printf("\n\n")

	// Soal 2
	smartphone := utils.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(smartphone.Display())
	fmt.Printf("\n\n")

	// Soal 3
	fmt.Println(utils.LuasPersegi(4, true))
	fmt.Println(utils.LuasPersegi(8, false))
	fmt.Println(utils.LuasPersegi(0, true))
	fmt.Println(utils.LuasPersegi(0, false))
	fmt.Printf("\n\n")

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefixString := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := utils.SumNumbers(angkaPertama...)
	total += utils.SumNumbers(angkaKedua...)

	fmt.Printf("%s%d + %d + %d + %d = %d\n", prefixString, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
}
