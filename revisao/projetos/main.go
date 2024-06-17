package main

import "fmt"

func main() {
	menu := map[string]float64{
		"pizza": 40.00,
		"suco": 12.50,
	}
	for k, v := range menu {
		fmt.Println(k, ": R$", v)
	}

	contaNova := novaConta("astrubal")
	fmt.Println(contaNova)
}
