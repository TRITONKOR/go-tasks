//go:build task46

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

	fmt.Print("Введіть val1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	val1, err1 := strconv.ParseFloat(input1, 64)
	if err1 != nil {
		fmt.Printf("'%s' не є числовим значенням\n", input1)
		return
	}

	fmt.Print("Введіть val2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	val2, err2 := strconv.ParseFloat(input2, 64)
	if err2 != nil {
		fmt.Printf("'%s' не є числовим значенням\n", input2)
		return
	}

	result := (val1 * 3 + val1) / 4 - val2

	fmt.Printf("Результат: %g\n", result)
}
