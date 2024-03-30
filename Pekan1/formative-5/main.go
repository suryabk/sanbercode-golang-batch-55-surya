package main

import (
	"fmt"
	"strings"
)

func main() {

	// jawaban Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// jawaban Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// jawaban soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	// jawaban soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(film, duration, genre, year string) {
		newFilm := map[string]string{"genre": genre, "jam": duration, "tahun": year, "title": film}
		dataFilm = append(dataFilm, newFilm)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}

func buahFavorit(name string, buah ...string) (sentence string) {
	fruits := `"` + strings.Join(buah, `", "`) + `"`
	sentence = fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", name, fruits)
	return
}

func introduce(name, gender, job, age string) (sentence string) {
	sapaan := "Pak"
	if gender == "perempuan" {
		sapaan = "Bu"
	}
	sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", sapaan, name, job, age)
	return
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

func kelilingPersegiPanjang(panjang int, lebar int) (hasil int) {
	hasil = 2 * (panjang + lebar)
	return
}

func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}
