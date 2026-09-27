//go:build task52

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

	fmt.Print("Введіть ціле число: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Printf("'%s' не є числом\n", num)
		return
	}


	result := (math.Sqrt(num + math.Sqrt(math.Pow(num, num)))) / 7

	fmt.Printf("Результат: %g\n", result)
}
