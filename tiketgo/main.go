package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Mengatur penanganan file statis(biar folder static terbaca)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Menentukan penanganan untuk halaman utama
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parsing templat HTML(nantine akan dieksekusi ning ngisor)
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println("Error parsing template:", err)
			return
		}

		// Menyajikan templat HTML(mengeksekusi index.html)
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println("Error executing template:", err)
			return
		}
	})

	// Menjalankan server di port 8082(wis ngerti lah yaa)
	fmt.Println("Server berhasil dijalankan di http://localhost:8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
