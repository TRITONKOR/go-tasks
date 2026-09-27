//go:build task36

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readIntSlice(reader *bufio.Reader, prompt string) ([]int, error) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fields := strings.Fields(input)

	numbers := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func readStringSlice(reader *bufio.Reader, prompt string) []string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strings.Fields(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	keys, errKeys := readIntSlice(reader, "Введіть список ключів через пробіл: ")
	if errKeys != nil {
		fmt.Println("Ключі мають бути цілими числами!")
		return
	}

	values := readStringSlice(reader, "Введіть список значень через пробіл: ")

	if len(keys) != len(values) {
		fmt.Println("Кількість ключів та значень має бути однаковою!")
		return
	}

	myMap := make(map[int]string)
	for i := 0; i < len(keys); i++ {
		myMap[keys[i]] = values[i]
	}

	fmt.Print("Введіть значення A для видалення: ")
	inputA, _ := reader.ReadString('\n')
	targetA := strings.TrimSpace(inputA)

	for key, val := range myMap {
		if val == targetA {
			delete(myMap, key)
		}
	}

	fmt.Printf("Отримана map після видалення значення \"%s\": %v\n", targetA, myMap)
}
