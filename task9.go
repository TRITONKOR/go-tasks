//go:build task9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть str1: ")
	str1, _ := reader.ReadString('\n')

	fmt.Print("Введіть str2: ")
	str2, _ := reader.ReadString('\n')

	word1 := strings.TrimSpace(str1)
	word2 := strings.TrimSpace(str2)


	runes1 := []rune(word1)
	runes2 := []rune(word2)

	if len(runes1) == 0 || len(runes2) == 0 {
		fmt.Println("Помилка: обидва рядки мають бути не порожніми!")
		return
	}

	first1 := runes1[0]
	middle1 := runes1[len(runes1)/2]
	last1 := runes1[len(runes1)-1]

	first2 := runes2[0]
	middle2 := runes2[len(runes2)/2]
	last2 := runes2[len(runes2)-1]

	str3 := string([]rune{first1, middle1, last1, first2, middle2, last2})

	fmt.Printf("Результат: %s\n", str3)
}
