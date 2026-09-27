//go:build task24

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть зріз чисел через пробіл: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	rawNumbers := strings.Fields(input)

	numbers := make([]int, 0, len(rawNumbers))

	for _, raw := range rawNumbers {
		num, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			fmt.Printf("Помилка: '%s' не є дійсним числом!\n", raw)
			return
		}
		rounded := int(math.Floor(num))
		numbers = append(numbers, rounded)
	}

	fmt.Printf("Результат: %v\n", numbers)

}
