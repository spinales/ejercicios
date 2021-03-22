package main

import (
	"html/template"
	"log"
	"net/http"
)

func server() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./files"))))

	http.HandleFunc("/form", form)
	http.HandleFunc("/process", formProcess)
	http.HandleFunc("/home", home)

	log.Println("app runing :8000")
	http.ListenAndServe(":8000", nil)
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
