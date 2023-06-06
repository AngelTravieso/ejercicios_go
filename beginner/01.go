package beginner

import "fmt"

func Ex001() {

	fmt.Println("Numeros divisibles por 7 que no son multiplos de 5:")

	numberRange(2000, 3200)

}

func numberRange(low int, hight int) {
	for i := low; i <= hight; i++ {
		if i%7 == 0 && i%5 != 0 {
			fmt.Printf("%d,", i)
		}
	}
}
