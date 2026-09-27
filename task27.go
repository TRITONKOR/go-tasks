//go:build task27

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть зріз чисел через пробіл: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	slice := strings.Fields(input)

	length := len(slice)
	if length == 0 {
		fmt.Println("Ви ввели порожній зріз!")
		return
	}

	fmt.Print("Введіть початковий індекс a: ")
	inputA, _ := reader.ReadString('\n')
	a, errA := strconv.Atoi(strings.TrimSpace(inputA))

	fmt.Print("Введіть кінцевий індекс b: ")
	inputB, _ := reader.ReadString('\n')
	b, errB := strconv.Atoi(strings.TrimSpace(inputB))


	if errA != nil || errB != nil {
		fmt.Println("Індекси a та b мають бути цілими числами!")
		return
	}

	if a < 0 || b >= length || a > b {
		fmt.Printf("Некоректний діапазон. Індекси мають бути від 0 до %d, і a <= b.\n", length-1)
		return
	}

	subSlice := slice[a : b+1]

	fmt.Printf("Результат: %v\n", subSlice)

}
