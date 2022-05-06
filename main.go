package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method  not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Post Request Successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name - %v\n", name)
	fmt.Fprintf(w, "address - %v\n", address)

}

func main() {
	// http fileserver helps reference a file or folder in your app
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// handleFunc tells a route what function to use
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting cat server/n")
	// this starts the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
