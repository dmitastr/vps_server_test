package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	address := ":" + port

	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	useName := name
	if name == "" {
		useName = "World"
	}
	log.Printf("Receive request with name=%s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello %s. This is test server", useName)))
}
