package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// When the form is filled in these variables will be taken in
	fmt.Fprintf(w, "Post request sucessful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// w = response writer, what server sends r= request, what user sends  the *(star) is the pointer
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// If URL Path !=(is not) Hello then send 404
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// If user attempds get on Hello
	if r.Method != "Get" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return

	}
	fmt.Fprintf(w, "hello!")

}
func main() {
	// Fileserver variable links to static folder
	fileServer := http.FileServer(http.Dir("./static"))

	// handles the Root "/" Route
	http.Handle("/", fileServer)

	// handles the form and will link us to the Form
	http.HandleFunc("/form", formHandler)

	// Will just print Hello as there is no Form
	http.HandleFunc("/hello", helloHandler)

	// Terminal Message for Port 8000
	fmt.Printf("Starting Server at Port 8080\n")
	// Sets listener to Port 8000 also creates server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// returns Error if cannot connect. Uses Log Import Package
		log.Fatal(err)
	}
}
