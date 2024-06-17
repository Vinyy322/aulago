package main

import "fmt"

func main() {
	numbers := [3]int{1, 2, 3}

	for k, v := range numbers {
		fmt.Println(k, "_", v)
	}

	var frutas = [3]string{"uva", "maça", "pera"}

	for k, v := range frutas {
		fmt.Println(k, "_", v)
	}

	matriz := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matriz[0][0])

	precos := [4]float64{10.5, 20.75, 30.0, 40.25}

	fmt.Println(precos[0])
}
