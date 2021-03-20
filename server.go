package main

import (
	"fmt"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Server basico")
	})

	// http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))

	// http.HandleFunc("/about", func(rw http.ResponseWriter, r *http.Request) {
	// 	lp := filepath.Join("files", "index.html")
	// 	fp := filepath.Join("files", filepath.Clean(r.URL.Path))

	// 	tmpl, _ := template.ParseFiles(lp, fp)
	// 	tmpl.ExecuteTemplate(rw, "index", nil)
	// })

	log.Println("app runing :8000")
	http.ListenAndServe(":8000", nil)
}
