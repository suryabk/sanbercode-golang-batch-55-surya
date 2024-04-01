package main

import "fmt"

func main() {
	// Soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64
	hitungLingkaran(&luasLingkaran, &kelilingLingkaran, 5)
	fmt.Printf("Luas Lingkaran : %.2f\n", luasLingkaran)
	fmt.Printf("keliling Lingkaran : %.2f\n", kelilingLingkaran)
	fmt.Printf("\n\n")

	// Soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
	fmt.Printf("\n\n")

	// Soal 3
	var buah = []string{}
	buahFavorit(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, fruit := range buah {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}
	fmt.Printf("\n\n")

	// Soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data
	for index, item := range dataFilm {
		fmt.Printf("%d. title : %s\n   duration : %s\n   genre : %s\n   year : %s\n\n", index+1, item["title"], item["duration"], item["genre"], item["year"])
	}
}

func tambahDataFilm(title, duration, genre, year string, data *[]map[string]string) {
	*data = append(*data, map[string]string{"genre": genre, "duration": duration, "year": year, "title": title})
}

func buahFavorit(arr *[]string, input ...string) {
	*arr = append(*arr, input...)
}

func introduce(sentence *string, name, gender, job, age string) {
	sapaan := "Pak"
	if gender == "perempuan" {
		sapaan = "Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", sapaan, name, job, age)
	return
}

func hitungLingkaran(luasLingkaran, kelilingLingkaran *float64, radius float64) {
	var phi float64 = 3.14
	*luasLingkaran = phi * radius * radius
	*kelilingLingkaran = phi * radius * 2
	return
}
