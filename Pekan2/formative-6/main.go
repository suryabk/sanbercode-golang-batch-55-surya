package main

import "fmt"

type Hasil struct{ luas, keliling float64 }

func main() {
	// Soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64
	HitungLingkatan(&luasLingkaran, &kelilingLingkaran, 5)
	fmt.Println(luasLingkaran, kelilingLingkaran)

	// Soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
}

func introduce(sentence *string, name, gender, job, age string) {
	sapaan := "Pak"
	if gender == "perempuan" {
		sapaan = "Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", sapaan, name, job, age)
	return
}

func HitungLingkatan(luasLingkaran, kelilingLingkaran *float64, radius float64) {
	var phi float64 = 3.14
	*luasLingkaran = phi * radius * radius
	*kelilingLingkaran = phi * radius * 2
	return
}
