package main

import (
	"github.com/gorilla/websocket"
	. "github.com/ilya-shikhaleev/drum-orchestra/lib"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	ADDR string = ":8083"
)

var pool *Pool

func homeHandler(c http.ResponseWriter, r *http.Request) {
	var homeTemplate = template.Must(template.ParseFiles("web/index.html"))
	data := struct {
		Title string
	}{"Добро пожаловать"}
	homeTemplate.Execute(c, data)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//1024 - buffer size
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", http.StatusBadRequest)
		return
	} else if err != nil {
		return
	}
	NewConnection(ws, pool)
	log.Println("New connection created")
}

func main() {
	log.Println("Start our app")
	pool = NewPool()
	log.Println("Connection pool created")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)

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
