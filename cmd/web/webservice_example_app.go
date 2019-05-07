package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/fact", factHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func factHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Input request with headers: %v\n", r.Header)
	fmt.Fprint(w, "Rock is not dead")
}
