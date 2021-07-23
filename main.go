package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Main function called")
	handleRequests()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequests() {
	http.HandleFunc("/", homePage)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
