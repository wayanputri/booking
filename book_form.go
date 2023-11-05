package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type DataForm struct {
	Name     string
	Email    string
	Phone    string
	Address  string
	Location string
	Guests   int
	Arrivals string
	Leaving  string
}

var templates = template.Must(template.ParseGlob("*.html"))

func main() {
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/submit", submitHandler)
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "book.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		data := DataForm{
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Address:  r.FormValue("address"),
			Location: r.FormValue("location"),
			Guests:   0, // Ganti dengan parsing angka sesuai kebutuhan Anda
			Arrivals: r.FormValue("arrivals"),
			Leaving:  r.FormValue("leaving"),
		}
		// Lakukan sesuatu dengan data yang Anda terima, misalnya menyimpannya ke database
		// Atau tampilkan respon sukses
		fmt.Fprintf(w, "Data telah diterima: %+v", data)
	} else {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
