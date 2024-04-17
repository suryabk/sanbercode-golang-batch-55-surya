package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func makeIndeks(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	} else {
		return "E"
	}
}

// jawaban no 2
func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhsNilai := nilaiNilaiMahasiswa
		dataMovies, err := json.Marshal(mhsNilai)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

// Jawaban no 1
func PostNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mhs NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			id := len(nilaiNilaiMahasiswa) + 1
			nama := r.PostFormValue("Nama")
			mataKuliah := r.PostFormValue("MataKuliah")
			getNilai := r.PostFormValue("Nilai")
			nilai, _ := strconv.Atoi(getNilai)
			indeks := makeIndeks(nilai)
			Mhs = NilaiMahasiswa{ID: uint(id), Nama: nama, MataKuliah: mataKuliah, Nilai: uint(nilai), IndeksNilai: indeks}
		}
		// filter agar nilai yg diinput hanya antar 0-100
		if Mhs.Nilai > 100 || Mhs.Nilai < 0 {
			http.Error(w, "Nilai minimal adalah 0, maksimal adalah 100", http.StatusBadRequest)
			return
		}
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Mhs)
		dataNilai, _ := json.Marshal(Mhs) // to byte
		w.Write(dataNilai)                // cetak di browser
		return

	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func main() {
	http.HandleFunc("/nilai", getNilaiMahasiswa)
	http.Handle("/post_nilai", Auth(http.HandlerFunc(PostNilaiMahasiswa)))
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
