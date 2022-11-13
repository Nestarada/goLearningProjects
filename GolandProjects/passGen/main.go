package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func generator() {
	var size int
	var pass []string
	fmt.Printf("Enter a password length: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		panic("error during entering password length")
	}
	for i := 0; i < size; i++ {
		pass = append(pass, string(rand.Intn(127-33)+33))
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	htmlPage, _ := template.ParseFiles("index.html")
	err := htmlPage.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		panic("error during parsing html")
	}
}

func main() {
	generator()
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	err := http.ListenAndServe("192.168.1.100:80", mux)
	if err != nil {
		panic("error during server listening")
	}
}
