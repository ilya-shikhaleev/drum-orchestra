package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	ADDR string = ":8083"
)

func homeHandler(c http.ResponseWriter, r *http.Request) {
	var homeTemplate = template.Must(template.ParseFiles("web/index.html"))
	data := struct {
		Title string
	}{"Добро пожаловать"}
	homeTemplate.Execute(c, data)
}

func main() {
	log.Println("Start our app")

	http.HandleFunc("/", homeHandler)

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./web/img/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css/"))))
	http.Handle("/audio/", http.StripPrefix("/audio/", http.FileServer(http.Dir("./web/audio/"))))

	s := &http.Server{
		Addr:           ADDR,
		Handler:        nil,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
