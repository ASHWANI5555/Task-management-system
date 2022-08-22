package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.html", nil)
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", fs))
	//r.Handle("/assets/", http.StripPrefix("/assets", fs))
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":7000", r)
}
