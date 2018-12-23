package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(1, 2))
}

// add two numbers
func Sum(a, b int) int {
	return a + b
}

// sub two numbers
func Sub(a, b int) int {
	return a - b
}

// TODO 1 Добавить метод вида num.add()
// TODO 2 Добавить ссетевой сервер
// TODO 3 Добавить обработчик / - прибавить один к счётчику посещений
