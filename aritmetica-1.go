package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Ingrese expresión aritmetica con un único operador")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	linea := scanner.Text()
	fmt.Println("Ingresado:", linea)
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(" Me aseguro el ingreso", r)
	//		scanner := bufio.NewScanner(os.Stdin)
	//		scanner.Scan()
	//		linea = scanner.Text()
	//	}
	//}()
	// variable para guardar el número de operadores aritméticos
	// podría ir a una goroutine
	var qty int
	suma := strings.Contains(linea, "+")
	res := strings.Contains(linea, "-")
	prod := strings.Contains(linea, "*")
	division := strings.Contains(linea, "/")
	qty = 0
	if suma {
		qty = qty + 1
	}
	if res {
		qty = qty + 1
	}
	if prod {
		qty = qty + 1
	}
	if division {
		qty = qty + 1
	}
	fmt.Println("Cantidad de operadores:", qty)
	if qty == 0 {
		fmt.Println("sin operadores, fin...")
		os.Exit(0)
	}

	//	os.Exit(0)
	// deseo localizar la posición de los operadores

	var plus, rest, produ, divi int = 1000, 1000, 1000, 1000

	plus = strings.Index(linea, "+")
	rest = strings.Index(linea, "-")
	produ = strings.Index(linea, "*")
	divi = strings.Index(linea, "/")

	// del índice, los valores en -1 debo considerarlos valores  mayores.
	if plus == -1 {
		plus = 1000
	}
	if rest == -1 {
		rest = 1000
	}
	if divi == -1 {
		divi = 1000
	}
	if produ == -1 {
		produ = 1000
	}

	fmt.Println("Posicion de operadores Suma, Resta, Producto, Division: ", plus, rest, produ, divi)

	//}
	// debo organizar de menor a mayor las operaciones según el indice encontrado

	if plus < rest && plus < produ && plus < divi {

		valores := strings.Split(linea, "+")

		//parte1, _ := strconv.Atoi(valores[0])
		parte1, err := strconv.Atoi(valores[0])
		fmt.Println("Valor parte1 suma:", parte1)
		fmt.Println(valores[0], valores[1])
		if err != nil {
			panic("Recuperando error suma parte1")
		}
		parte2, err := strconv.Atoi(valores[1])
		fmt.Println("Valor parte2 suma:", parte2)
		if err != nil {
			panic("Recuperando error suma parte2")
		}
		fmt.Println("Resultado:", parte1+parte2)
	}
	if rest < plus && rest < produ && rest < divi {
		valores := strings.Split(linea, "-")
		qty = qty + 1
		//parte1, _ := strconv.Atoi(valores[0])
		parte1, err := strconv.Atoi(valores[0])
		if err != nil {
			panic("Recuperando error resta parte1")
		}
		parte2, err := strconv.Atoi(valores[1])
		if err != nil {
			panic("Recuperando error resta parte2")
		}
		fmt.Println("Resultado:", parte1-parte2)
	}
	if produ < rest && produ < plus && produ < divi {
		valores := strings.Split(linea, "*")
		//parte1, _ := strconv.Atoi(valores[0])
		parte1, err := strconv.Atoi(valores[0])
		if err != nil {
			panic("Recuperando error producto parte1")
		}
		parte2, err := strconv.Atoi(valores[1])
		if err != nil {
			panic("Recuperando error producto parte2")
		}
		fmt.Println("Resultado:", parte1*parte2)
	}
	if divi < plus && divi < produ && divi < rest {
		valores := strings.Split(linea, "/")

		//parte1, _ := strconv.Atoi(valores[0])
		parte1, err := strconv.Atoi(valores[0])
		if err != nil {
			panic("Recuperando error division parte1")
		}
		parte2, err := strconv.Atoi(valores[1])
		if err != nil {
			panic("Recuperando error division parte2")
		}
		fmt.Println("Resultado:", parte1/parte2)
	}

}
