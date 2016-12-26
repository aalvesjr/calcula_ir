package main

import (
	"fmt"
	salario "github.com/aalvesjr/salario"
	"log"
	"os"
	"strconv"
)

var (
	ajuda string = `
	São necessários dois argumentos para executar esse script: salário e descontos
	Para calcular o IR de um salário de 5.000,00 e com descontos de 200,00, execute:

	./calcula_ir 5000.00 200.00
`
)

func main() {
	// Recebe o salário e descontos por argumentos
	// ./calcula_ir 6120.32 501.32
	if len(os.Args) < 3 {
		log.Fatal(ajuda)
	}
	valor, err := strconv.ParseFloat(os.Args[1], 32)
	descontos, err2 := strconv.ParseFloat(os.Args[2], 32)

	if err != nil || valor < 0 {
		log.Fatal("Formato de salário inválido! Tente: 12345.67")
	}
	if err2 != nil || descontos < 0 {
		log.Fatal("Formato de desconto inválido! Tente: 12345.67")
	}

	s := salario.NewSalario(float32(valor), float32(descontos))

	fmt.Printf("Salário Bruto   => R$ %.2f\n", s.Bruto)
	fmt.Printf("Descontos       => R$ %.2f\n", s.Descontos)
	fmt.Println("------------INSS-------------")
	fmt.Printf("Base INSS       => R$ %.2f\n", s.BaseINSS)
	fmt.Printf("Aliquota INSS   => %.2f%%\n", s.AliquotaINSS)
	fmt.Printf("Valor INSS      => R$ %.2f\n", s.INSS)
	fmt.Println("-------------IR--------------")
	fmt.Printf("Base IR         => R$ %.2f\n", s.BaseIR)
	fmt.Printf("Aliquota IR     => %.2f%%\n", s.AliquotaIR)
	fmt.Printf("IR sem desconto => R$ %.2f\n", s.IRSemDesconto)
	fmt.Printf("Desconto do IR  => R$ %.2f\n", s.DescontoIR)
	fmt.Printf("Valor IR        => R$ %.2f\n", s.IR)
	fmt.Println("-----------------------------")
	fmt.Printf("Salário Liquido => R$ %.2f\n", s.Liquido)
}
