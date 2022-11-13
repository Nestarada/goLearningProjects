package main

import (
	"html/template"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	htmlPage, _ := template.ParseFiles("templates/index.html")
	htmlPage.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	mux := http.NewServeMux()
	assets := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assets))
	mux.HandleFunc("/", mainPage)
	http.ListenAndServe("192.168.1.100:80", mux)
}
