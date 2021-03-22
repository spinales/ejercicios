package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func server() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./files"))))

	http.HandleFunc("/secret", secret)
	http.HandleFunc("/form", form)
	http.HandleFunc("/process", formProcess)
	http.HandleFunc("/home", home)

	log.Println("app runing :8000")
	http.ListenAndServe(":8000", nil)
}

func secret(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()

	if !ok {
		fmt.Println("Error parsing basic auth")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if u != "admin" {
		fmt.Println("Username provided is incorrect: ", u)
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	if p != "Pass123" {
		fmt.Println("Password provided is incorrect: ", p)
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	d := struct {
		Tittle  string
		Content string
	}{
		"Pagina secreta",
		"contenido de pagina secreta",
	}

	tpl := template.Must(template.ParseGlob("files/*.html"))
	tpl.ExecuteTemplate(w, "default.html", d)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("files/*.html"))
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func form(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("files/*.html"))
	tpl.ExecuteTemplate(w, "form.html", nil)
}

func formProcess(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("files/*.html"))

	if r.Method != "GET" {
		http.Redirect(w, r, "/form", http.StatusSeeOther)
	}

	data := struct {
		Firstname string
		Lastname  string
	}{
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
	}

	tpl.ExecuteTemplate(w, "display.html", data)
}
