//go:build task34

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	keys, err1 := readIntSlice(reader, "Введіть список ключів через пробіл: ")
	values, err2 := readIntSlice(reader, "Введіть список значень через пробіл: ")

	if err1 != nil || err2 != nil {
		fmt.Println("Елементи списків мають бути цілими числами!")
		return
	}

	if len(keys) != len(values) {
		fmt.Println("Кількість ключів та значень має бути однаковою!")
		return
	}

	myMap := make(map[int]int)
	for i := 0; i < len(keys); i++ {
		myMap[keys[i]] = values[i]
	}

	fmt.Print("Введіть шуканий ключ A: ")
	inputA, _ := reader.ReadString('\n')
	keyA, errA := strconv.Atoi(strings.TrimSpace(inputA))

	fmt.Print("Введіть шукане значення B: ")
	inputB, _ := reader.ReadString('\n')
	valB, errB := strconv.Atoi(strings.TrimSpace(inputB))

	if errA != nil || errB != nil {
		fmt.Println("A та B мають бути цілими числами!")
		return
	}

	val, exists := myMap[keyA]

	fmt.Printf("\nСформована map: %v\n", myMap)

	if exists && val == valB {
		fmt.Printf("Результат: Ключ %d існує і його значення дорівнює %d.\n", keyA, valB)
	} else {
		fmt.Printf("Результат: Ключ %d відсутній у словнику.\n", keyA)
	}
}
