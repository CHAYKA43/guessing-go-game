package main

import (
	"fmt"
	"math/rand"
)

// единственное что я украл у нейросети
func randomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	// Инициализация и настройка игры
	range1 := 1
	range2 := 20
	number := randomRange(range1, range2)
	attempts := 10
	var input_number int

	// Главный цикл игры
	fmt.Printf("Привет! Я загадал число от %v до %v, у тебя есть %v попыток чтобы его угадать :D\n", range1, range2, attempts)
	for input_number != number && attempts > 0 {
		fmt.Print("Введите число: ")
		_, err := fmt.Scanln(&input_number)
		attempts--
		if err != nil {
			fmt.Println("Введите число!!")
			break
		} else if number < input_number {
			fmt.Printf("Число меньше! Осталось попыток: %v\n", attempts)
		} else if number > input_number {
			fmt.Printf("Число больше! Осталось попыток: %v\n", attempts)
		}
	}
	if input_number == number {
		fmt.Println("Молодец!")
	} else if attempts <= 0 {
		fmt.Printf("Попытки кончились! Мое число было: %v\n", number)
	}
}
