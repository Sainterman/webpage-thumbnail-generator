package main

import (
	"fmt"
	"net/http"
	"log"
)
func main() {
	http.HandleFunc("/", homePageHandler)
	fmt.Println("Server listening on port 5000")
	log.Panic(
		http.ListenAndServe(":5000", nil),
	)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}