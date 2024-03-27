package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringToInt(word string) int {
	var num, err = strconv.Atoi(word)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return num
}

func arraySum(arr []int) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

func main() {
	// Soal 1
	var text1 string
	text1 = "Bootcamp "
	var text2 string
	text2 = "Digital "
	var text3 string = "Skill "
	var text4 string = "Sanbercode "
	var text5 string = "Golang"

	fmt.Println(text1 + text2 + text3 + text4 + text5)

	// Soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", text5, 1)

	fmt.Println(halo)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kalimatGabungan := []string{strings.Title(kataPertama), strings.Title(kataKedua), strings.Replace(kataKetiga, "r", "R", 1), strings.ToUpper(kataKeempat)}

	fmt.Println(strings.Join(kalimatGabungan, " "))

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	arrayNum := []int{stringToInt(angkaPertama), stringToInt(angkaKedua), stringToInt(angkaKetiga), stringToInt(angkaKeempat)}

	fmt.Printf("angka: %d, jumlah: %d\n", arrayNum, arraySum(arrayNum))

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)
	fmt.Printf("%s - %d\n", kalimat, angka)

}
