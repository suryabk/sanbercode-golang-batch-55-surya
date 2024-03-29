package main

import (
	"fmt"
)

func main() {

	// Soal 1
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 0 {
			fmt.Printf("%d - Berkualitas\n", i)
		} else {
			fmt.Printf("%d - Santai\n", i)
		}
	}
	fmt.Println() //untuk buat baris baru

	// Soal 2
	for x := 1; x <= 7; x++ {
		star := ""
		for y := 1; y <= x; y++ {
			star += "#"
		}
		fmt.Printf("%s\n", star)
	}
	fmt.Println() //untuk buat baris baru

	// Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	var newKalimat = append(kalimat[:1], kalimat[3:]...)
	for _, kata := range newKalimat {
		fmt.Printf("%s ", kata)
	}
	fmt.Println() //untuk buat baris baru
	fmt.Println() //untuk buat baris baru

	// Soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}
	fmt.Println() //untuk buat baris baru

	// Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}
	fmt.Printf("Volume Balok = %d\n", satuan["panjang"]*satuan["lebar"]*satuan["tinggi"])
}
