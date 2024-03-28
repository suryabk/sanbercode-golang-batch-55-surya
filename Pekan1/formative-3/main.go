package main

import (
	"fmt"
	"strconv"
)

func stringToInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return intVal
}

func main() {

	// Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	luasPersegiPanjang = stringToInt(panjangPersegiPanjang) * stringToInt(lebarPersegiPanjang)
	fmt.Printf("luas persegi panjang = %d\n", luasPersegiPanjang)

	kelilingPersegiPanjang = 2 * (stringToInt(panjangPersegiPanjang) + stringToInt(lebarPersegiPanjang))
	fmt.Printf("keliling persegi panjang = %d\n", kelilingPersegiPanjang)

	luasSegitiga = stringToInt(alasSegitiga) * stringToInt(tinggiSegitiga) * 1 / 2
	fmt.Printf("luas segitiga = %d\n", luasSegitiga)

	// Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Index nilaiJohn adalah A")
	} else if nilaiJohn >= 70 {
		fmt.Println("Index nilaiJohn adalah B")
	} else if nilaiJohn >= 60 {
		fmt.Println("Index nilaiJohn adalah C")
	} else if nilaiJohn >= 50 {
		fmt.Println("Index nilaiJohn adalah D")
	} else {
		fmt.Println("Index nilaiJohn adalah E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Index nilaiDoe adalah A")
	} else if nilaiDoe >= 70 {
		fmt.Println("Index nilaiDoe adalah B")
	} else if nilaiDoe >= 60 {
		fmt.Println("Index nilaiDoe adalah C")
	} else if nilaiDoe >= 50 {
		fmt.Println("Index nilaiDoe adalah D")
	} else {
		fmt.Println("Index nilaiDoe adalah E")
	}

	// Soal 3
	var tanggal = 17
	var bulan = 8
	var tahun = 1945
	var namaBulan string

	tanggal = 27
	bulan = 3
	tahun = 2000

	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	}

	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	// Soal 4
	var tahunLahir int = 2000
	if tahunLahir >= 1995 {
		fmt.Printf("Generasi Z")
	} else if tahunLahir >= 1980 {
		fmt.Printf("Generasi Y (Millenials)")
	} else if tahunLahir >= 1965 {
		fmt.Printf("Generasi X")
	} else if tahunLahir <= 1964 && tahunLahir >= 1944 {
		fmt.Printf("Baby boomer")
	}
}
