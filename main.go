package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Main function called")
	handleRequests()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", helloPage)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func helloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HelloPage!")
	fmt.Println("Endpoint Hit: homePage")
}
