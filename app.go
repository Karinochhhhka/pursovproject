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
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", searchHandler)
	http.ListenAndServe(":8080", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем введенное слово из параметра запроса или из формы ввода
	query := r.FormValue("query")

	// Текст, в котором будет осуществляться поиск
	text := "Это пример текста, в котором мы будем искать введенное слово."

	// Осуществляем поиск в тексте
	found := strings.Contains(text, query)

	// Выводим HTML форму с текстом, полем ввода и кнопкой
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Поиск слова</title>
		</head>
		<body>
			<h1>Поиск слова в тексте</h1>
			<form method="post">
				<label for="query">Введите слово:</label>
				<input type="text" id="query" name="query">
				<input type="submit" value="Поиск">
			</form>
			<hr>
			<p>Текст: %s</p>
			<p>Введенное слово: %s</p>
	`, text, query)

	// Выводим результат поиска на страницу
	if found {
		fmt.Fprintf(w, "<p>Слово найдено в тексте.</p>")
	} else {
		fmt.Fprintf(w, "<p>Слово не найдено в тексте.</p>")
	}

	fmt.Fprintf(w, `
		</body>
		</html>
	`)
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




