package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called from github")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
