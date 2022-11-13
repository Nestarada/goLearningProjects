package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func cfgCreator(w http.ResponseWriter, r *http.Request) {
	reqSize, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	buff := make([]byte, reqSize)
	r.Body.Read(buff)
	fmt.Println(string(buff))
	//os.WriteFile("config/"+req[3]+".json", []byte(""), 0666)
	//fmt.Println("File created successfully")
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	htmlPage := template.Must(template.New("").ParseFiles("html/index.html"))
	err := htmlPage.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		panic("error while executing template")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		mainPage(w, r)
	case "/create":
		cfgCreator(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", mux)
}
