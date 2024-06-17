package main

import "fmt"

func socoro(){
	fmt.Print("não XD")
}

func main(){
	idade := 39
	fmt.Println(idade <= 40)
	fmt.Println(idade >= 40)
	fmt.Println(idade == 40)
	fmt.Println(idade != 40)

	if(idade > 40){
		fmt.Print("você é idoso fudido\n")
	}else if(idade < 40){
		fmt.Print("você não é macho\n")
	}
	socoro()
}