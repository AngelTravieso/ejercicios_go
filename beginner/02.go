package beginner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number int
var err error

func Ex002() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("¡Factorial de un numero!")
	fmt.Print("Ingrese un numero: ")

	if scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es invalido " + err.Error())
		}

		fmt.Printf("Factorial de %d: %d", number, getFactorial(number))
	}

}

func getFactorial(number int) int {

	var result int = 1

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result

}
