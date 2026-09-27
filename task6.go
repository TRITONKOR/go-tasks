//go:build task6

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть слово: ")
	input, _ := reader.ReadString('\n')

	word := strings.TrimSpace(input)

	runes := []rune(word)
	length := len(runes)

	if length == 0 {
		fmt.Println("Ви нічого не ввели!")
		return
	}

	if length < 3  {
		fmt.Println("Слово має містити мінімум 3 символа")
		return
	}

	first := runes[(length/2) - 1]
	middle := runes[length/2]
	last := runes[(length/2) + 1]

	result := string([]rune{first, middle, last})

	fmt.Printf("Результат: %s\n", result)
}
