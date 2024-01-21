package main

import (
	"html/template"
	"net/http"
)
 
func main() {
	// Mengatur penanganan file statis
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Menentukan penanganan untuk halaman utama
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parsing templat HTML
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Menyajikan templat HTML
		tmpl.Execute(w, nil)
	})

	// Menjalankan server di port 8080
	http.ListenAndServe(":8082", nil)
}
