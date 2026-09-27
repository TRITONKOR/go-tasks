//go:build task5

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
		fmt.Println("Ви нічого не ввели")
		return
	}

	first := runes[0]
	middle := runes[length/2]
	last := runes[length-1]

	result := string([]rune{first, middle, last})

	fmt.Printf("Результат: %s\n", result)
}
