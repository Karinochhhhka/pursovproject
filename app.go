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
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func generateRandomNumber() int {
	return rand.Intn(100)
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		randomNumber := generateRandomNumber()
		text := fmt.Sprintf("Random number: %d", randomNumber)
		fmt.Fprintf(w, text)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Random Number Generator</title>
		</head>
		<body>
			<h1>Random Number Generator</h1>
			<form action="/" method="post">
				<button type="submit">Generate Random Number</button>
			</form>
			<div id="result"></div>
			<script>
				document.querySelector("form").addEventListener("submit", function(event) {
					event.preventDefault();
					fetch("/", { method: "POST" })
						.then(response => response.text())
						.then(text => {
							document.getElementById("result").textContent = text;
						});
				});
			</script>
		</body>
		</html>
	`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/number", numberHandler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// )

// func main() {
// 	http.HandleFunc("/", randomHandler)
// 	http.ListenAndServe(":8080", nil)
// }

// func randomHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		// Генерируем случайное число
// 		randomNumber := rand.Intn(100)

// 		// Преобразуем число в строку
// 		randomString := strconv.Itoa(randomNumber)

// 		// Выводим HTML с кнопкой и случайным числом
// 		fmt.Fprintf(w, `
// 			<!DOCTYPE html>
// 			<html>
// 			<head>
// 				<title>Случайное число</title>
// 			</head>
// 			<body>
// 				<h1>Случайное число</h1>
// 				<p>Сгенерированное число: %s</p>
// 				<form method="post">
// 					<input type="submit" value="Сгенерировать еще">
// 				</form>
// 			</body>
// 			</html>
// 		`, randomString)
// 	} else {
// 		// Выводим HTML с кнопкой для генерации случайного числа
// 		fmt.Fprintf(w, `
// 			<!DOCTYPE html>
// 			<html>
// 			<head>
// 				<title>Случайное число</title>
// 			</head>
// 			<body>
// 				<h1>Случайное число</h1>
// 				<form method="post">
// 					<input type="submit" value="Сгенерировать">
// 				</form>
// 			</body>
// 			</html>
// 		`)
// 	}
// }

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




