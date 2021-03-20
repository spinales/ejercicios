package main

import (
	"log"
	"net/http"
)

func server() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./files"))))

	log.Println("app runing :8000")
	http.ListenAndServe(":8000", nil)
}
