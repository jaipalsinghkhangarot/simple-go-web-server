package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "404 Not Found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(response, "hello!!")
}

func formHandler(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
		return
	}
	if request.Method != "POST" {
		http.Error(response, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(response, "POST request Successfull!!\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "Name = %s\n", name)
	fmt.Fprintf(response, "Address = %s\n", address)
}
