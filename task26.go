//go:build task26

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

	fmt.Print("Введіть зріз цілих чисел через пробіл: ")
	inputSlice, _ := reader.ReadString('\n')
	inputSlice = strings.TrimSpace(inputSlice)
	rawNumbers := strings.Fields(inputSlice)

	fmt.Print("Введіть число A: ")
	inputA, _ := reader.ReadString('\n')
	inputA = strings.TrimSpace(inputA)

	targetA, err := strconv.Atoi(inputA)
	if err != nil {
		fmt.Println("A має бути цілим числом")
		return
	}

	filteredSlice := make([]int, 0, len(rawNumbers))
	sum := 0

	for _, raw := range rawNumbers {
		num, err := strconv.Atoi(raw)
		if err != nil {
			fmt.Printf("'%s' не є цілим числом!\n", raw)
			return
		}

		if num != targetA {
			filteredSlice = append(filteredSlice, num)
			sum += num
		}
	}

	fmt.Printf("Зріз після видалення елементів зі значенням %d: %v\n", targetA, filteredSlice)
	fmt.Printf("Сума елементів, що залишилися: %d\n", sum)
}
