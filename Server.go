package main

import (
	"fmt" // пакет для форматированного ввода вывода
	"log"
	"net/http" // пакет для поддержки HTTP протокола
)

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	http.HandleFunc("/postform", EqualRouterHandler1)

	err := http.ListenAndServe(":8181", nil) // задаем слушать порт
	fmt.Println("Server is listening...")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Equal.html")
}

func EqualRouterHandler1(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	fmt.Fprintf(w, "Data is  %s", data)
}
