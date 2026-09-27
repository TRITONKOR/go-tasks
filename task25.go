//go:build task25

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть перший зріз чисел через пробіл: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	slice1 := strings.Fields(input1)

	fmt.Print("Введіть другий зріз чисел через пробіл: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	slice2 := strings.Fields(input2)

	slice3 := append(slice1, slice2...)

	fmt.Printf("Об'єднаний зріз: %v\n", slice3)

}
