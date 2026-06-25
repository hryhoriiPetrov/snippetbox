package main

import (
	"log"
	"net/http"
)

// http.ResponseWriter – дает параметры для сборки ответа и отправки юзеру
// *http.Request – указатель на структуру, которая содержит инфо о текущем запросе
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
