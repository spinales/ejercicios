package main

import (
	"log"
	"net/http"
)

func server() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./files"))))

	log.Println("app runing :4000")
	http.ListenAndServe(":4000", nil)
}
