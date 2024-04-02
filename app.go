package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
)

func generateRandomNumber() int {
     return rand.Intn(100)
}

func handler(w http.ResponseWriter, r *http.Request) {
    randomNumber := generateRandomNumber()
    fmt.Fprintf(w, "Random number: %d", randomNumber)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Running demo app. Press Ctrl+C to exit...")
    log.Fatal(http.ListenAndServe(":8888", nil))
}


// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, 世界")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Running demo app. Press Ctrl+C to exit...")
// 	log.Fatal(http.ListenAndServe(":8888", nil))
// }




