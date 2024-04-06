package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	// jawaban Soal 1
	defer showKalimat("Golang Backend Development", 2021)

	// jawaban Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
	fmt.Printf("\n\n")

	// jawaban Soal 3
	angka := 1

	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
	fmt.Printf("\n\n")

	// jawaban Soal 4
	var phones []string

	addPhones(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	sort.Strings(phones)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	i := 1
	for range ticker.C {
		if i <= len(phones) {
			fmt.Printf("%d. %s\n", i, phones[i-1])
			i++
		} else {
			break
		}
	}
	fmt.Printf("\n\n")

	// Soal 5

	// jawaban Soal 6
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)
	for value := range moviesChannel {
		fmt.Println(value)
	}
	fmt.Printf("\n\n")
}

func showKalimat(word string, year int) {
	fmt.Println("Jawaban No. 1")
	fmt.Println(word)
	fmt.Println(year)
}

func kelilingSegitigaSamaSisi(sisi int, description bool) string {
	if sisi == 0 {
		errMsg := "Maaf anda belum menginput sisi dari segitiga sama sisi\n"
		if description {
			return errMsg
		} else {
			defer endApp()
			panic(errMsg)
		}
	}
	keliling := sisi * 3
	if description {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm\n", sisi, keliling)
	} else {
		return fmt.Sprintf("%d\n", keliling)
	}
}

func endApp() {
	message := recover()
	fmt.Println("Panic Recover :", message)
}

func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Println("Jawaban no. 3")
	fmt.Println("Total :", *total)
	fmt.Printf("\n\n")
}

func getMovies(ch chan<- string, movies ...string) {
	for _, movie := range movies {
		ch <- movie
	}
	close(ch)
}

func addPhones(phones *[]string, data ...string) {
	*phones = append(*phones, data...)
}
