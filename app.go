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

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Отображение HTML-формы с кнопкой
		html := `
		<html>
		<head>
			<title>Random Number Generator</title>
		</head>
		<body>
			<button onclick="getRandomNumber()">Generate Random Number</button>
			<p id="randomNumber"></p>
			<script>
				function getRandomNumber() {
					fetch("/random")
						.then(response => response.text())
						.then(number => {
							document.getElementById("randomNumber").textContent = "Random number: " + number;
						})
						.catch(error => console.log(error));
				}
			</script>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html") // Установка заголовка Content-Type
		fmt.Fprint(w, html)
	} else if r.Method == "GET" && r.URL.Path == "/random" {
		// Генерация случайного числа и его отправка в ответе
		randomNumber := generateRandomNumber()
		fmt.Fprint(w, strconv.Itoa(randomNumber))
	}
}

func main() {
	http.HandleFunc("/", handler)
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




