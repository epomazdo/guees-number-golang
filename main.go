package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора
}

// RandomInt возвращает случайное число в диапазоне [min, max]
func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min // коррекция если min > max
	}
	return min + rand.Intn(max-min+1)
}

type Round struct {
	Difficulty string
	Number     int
	areWinned  bool
}

func (r *Round) StartGame() {
	switch {
	case r.Difficulty == "Easy":
		EasyDifficulty(r)
	case r.Difficulty == "Normal":
		NormalDifficulty(r)
	case r.Difficulty == "Hard":
		NormalDifficulty(r)
	}
}

func EasyDifficulty(r *Round) {
	r.Number = RandomInt(1, 10)
	fmt.Println("Лёгкий уровень: угадайте число от 1 до 10")

	GuessNumber(r)
}

func NormalDifficulty(r *Round) {
	r.Number = RandomInt(1, 25)
	fmt.Println("Обычный уровень: угадайте число от 1 до 25")

	GuessNumber(r)
}

func HardDifficulty(r *Round) {
	r.Number = RandomInt(1, 50)
	fmt.Println("Сложный уровень: угадайте число от 1 до 50")

	GuessNumber(r)
}

func GuessNumber(r *Round) {
	for i := 0; i < 10; i++ {
		var number int

		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		switch {
		case number < r.Number:
			fmt.Println("Секретное число больше")
		case number > r.Number:
			fmt.Println("Секретное число меньше")
		case number == r.Number:
			r.areWinned = true
		}

		if r.areWinned == true {
			fmt.Println("Вы угадали! Загаданное число равно", r.Number)
			break
		}
	}
	if r.areWinned == false {
		fmt.Println("10 попыток кончились - Вы проиграли.")
	}
}

func main() {
	fmt.Println("Игра 'Угадай число'")
	var diff int = 0

	for {
		if 1 <= diff && diff <= 3 {
			break
		}
		fmt.Println("Выберите сложность: \n 1 - легко \n 2 - обычно \n 3 - сложно")

		_, err := fmt.Scan(&diff)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
	}
	switch {
	case diff == 1:
		newRound := Round{"Normal", 0, false}
		newRound.StartGame()
	case diff == 2:
		newRound := Round{"Normal", 0, false}
		newRound.StartGame()
	case diff == 3:
		newRound := Round{"Normal", 0, false}
		newRound.StartGame()
	}
}
