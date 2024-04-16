// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"math/rand"
// )

// func generateRandomNumber() int {
//      return rand.Intn(100)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
//     randomNumber := generateRandomNumber()
//     fmt.Fprintf(w, "Random number: %d", randomNumber)
// }

// func main() {
//     http.HandleFunc("/", handler)
//     fmt.Println("Running demo app. Press Ctrl+C to exit...")
//     log.Fatal(http.ListenAndServe(":8888", nil))
// }


package main

import (
	"fmt"
	//"log"
	"net/http"
	//"math/rand"
	"strings"
)

func main() {
	http.HandleFunc("/", searchHandler)
	http.ListenAndServe(":8888", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем введенное слово из параметра запроса
	query := r.URL.Query().Get("query")

	// Текст, в котором будет осуществляться поиск
	text := "Это пример текста, в котором мы будем искать введенное слово."

	// Осуществляем поиск в тексте
	found := strings.Contains(text, query)

	// Выводим текст и результат на страницу
	fmt.Fprintf(w, "Текст: шрашврпшрпшкрпщугрпугпушпум слово слово слово %s\n", text)
	fmt.Fprintf(w, "Введенное слово: %s\n", query)
	if found {
		fmt.Fprintf(w, "Слово найдено в тексте.")
	} else {
		fmt.Fprintf(w, "Слово не найдено в тексте.")
	}
}


// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, 世界")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Running demo app. Press Ctrl+C to exit...")
// 	log.Fatal(http.ListenAndServe(":8888", nil))
// }








// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, 世界")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Running demo app. Press Ctrl+C to exit...")
// 	log.Fatal(http.ListenAndServe(":8888", nil))
// }




