package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

var users = map[string]string{
	"admin": "Pass123",
	"irvin": "123456",
}

func server() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./files"))))

	http.HandleFunc("/signin", signin)
	// authentication required
	http.HandleFunc("/secret", secret)
	// send basic form data
	http.HandleFunc("/form", form)
	// get data from form and print
	http.HandleFunc("/process", formProcess)
	http.HandleFunc("/home", home)

	log.Println("app runing :8000")
	http.ListenAndServe(":8000", nil)
}

func signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl := template.Must(template.ParseGlob("files/*.html"))
		tpl.ExecuteTemplate(w, "login.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	expectedPassword, ok := users[username]

	if !ok || expectedPassword != password {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	sessionToken, _ := uuid.NewV4()

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken.String(),
		Expires: time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "/secret", http.StatusPermanentRedirect)
}

func login(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("files/*.html"))
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func secret(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d := struct {
		Tittle  string
		Content string
	}{
		"Pagina secreta",
		fmt.Sprintf("Bienvenido %s a la pagina secreta", c.Value),
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
