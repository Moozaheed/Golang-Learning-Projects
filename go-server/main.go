package main

import (
    "fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	
	name := r.FormValue("name")
	message := r.FormValue("message")
	comment := r.FormValue("comment")

	fmt.Fprintf(w, "Name = %s\n", name) 
	fmt.Fprintf(w, "Message = %s\n", message)
	fmt.Fprintf(w, "Comment = %s\n", comment)

}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

    fmt.Fprintf(w, "Hello, World!")

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler) 
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("Server is listening port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
} 