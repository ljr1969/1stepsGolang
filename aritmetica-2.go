package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ExtractNumbersFromString extracts all numbers from a string and returns them as a slice of integers.
func ExtractNumbersFromString(s string) ([]int, error) {
	re := regexp.MustCompile(`\d+`) // Regular expression to match one or more digits
	matches := re.FindAllString(s, -1)

	numbers := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err // Return error if the matched string cannot be converted to an integer
		}
		numbers[i] = num
	}
	return numbers, nil
}
func ExtractOperadores(s string) (int, error) {
	var operadores int
	suma := strings.Contains(s, "+")
	res := strings.Contains(s, "-")
	prod := strings.Contains(s, "*")
	division := strings.Contains(s, "/")
	operadores = 0
	if suma {
		operadores = operadores + 1
	}
	if res {
		operadores = operadores + 1
	}
	if prod {
		operadores = operadores + 1
	}
	if division {
		operadores = operadores + 1
	}
	return operadores, nil
}
func removeNumbers(s string) string {
	re := regexp.MustCompile(`[0-9]+`)
	return re.ReplaceAllString(s, "")
}
func main() {
	fmt.Println("Ingrese expresión aritmetica como por ejemplo 100+30-50/10")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	linea := scanner.Text()
	// uso funcion para obtener valores numéricos de la entrada
	numbers, err := ExtractNumbersFromString(linea)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Numeros ingresados:", numbers)
	// obtener los operadores aritméticos
	operadores, err := ExtractOperadores(linea)
	fmt.Println("Operadores ingresados:", operadores)
	ordenopera := removeNumbers(linea)
	fmt.Println("Orden de los operadores ingresados:", ordenopera)
}
