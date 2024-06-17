package main

import "fmt"

type conta struct {
	nome    string
	itens   map[string]float64
	gorjeta float64
}

func novaConta(nome string) conta { // Corrigido o nome da função para novaConta
	c := conta{
		nome:    nome,
		itens:   map[string]float64{},
		gorjeta: 0.0,
	}
	return c
}

func (c conta) formatacao() conta {
	fs := "Detalhe da conta \n"
	var total float64 = 0

	for k,v := range c.itens{
		fs += fmt.Sprintf("%v....R$ %0.2f\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v....%0.2f","total",total)
	return fs
}