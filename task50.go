//go:build task50

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

	fmt.Print("Введіть числовий зріз mySlice: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	elements := strings.Fields(input)

	if len(elements) < 2 {
		fmt.Printf("Зріз повинен містити мінімум 2 числа\n")
		return
	}

	num1, err1 := strconv.ParseFloat(elements[0], 64)
	if err1 != nil {
		fmt.Printf("'%s' не є числом\n", num1)
		return
	}

	num2, err2 := strconv.ParseFloat(elements[len(elements)-1], 64)
	if err2 != nil {
		fmt.Printf("'%s' не є числом\n", num2)
		return
	}

	result := num1 + num2

	fmt.Printf("Результат: %g\n", result)
}
