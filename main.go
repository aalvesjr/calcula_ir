package main

import (
	"fmt"
	"github.com/aalvesjr/salary"
	"log"
	"os"
	"strconv"
)

var (
	help string = `
	Is necessary two arguments to execute this script: salary and discounts
	To calculate the IR from one salary of 5000.00 and with discounts of 200.00, execute:

	./calcula_ir 5000.00 200.00
`
)

func main() {
	// Receive the salary value and discounts from arguments
	// ./calcula_ir 6120.32 501.32
	if len(os.Args) < 3 {
		log.Fatal(help)
	}
	value, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil || value < 0 {
		log.Fatal("Wrong salary format! Try something like this: 12345.67")
	}

	discounts, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil || discounts < 0 {
		log.Fatal("Wrong discount format! Try something like this: 12345.67")
	}

	s := salary.NewSalary(float32(value), float32(discounts))

	fmt.Println("-------------Salary--------------")
	fmt.Printf("Salary Gross        => R$ %.2f\n", s.Gross)
	fmt.Printf("Discounts           => R$ %.2f\n", s.Discounts)
	fmt.Println("--------------INSS---------------")
	fmt.Printf("INSS Base           => R$ %.2f\n", s.INSSBase)
	fmt.Printf("INSS Rate           => %.2f%%\n", s.INSSRate*100)
	fmt.Printf("INSS Value          => R$ %.2f\n", s.INSS)
	fmt.Println("---------------IR----------------")
	fmt.Printf("IR Base             => R$ %.2f\n", s.IRBase)
	fmt.Printf("IR Rate             => %.2f%%\n", s.IRRate*100)
	fmt.Printf("IR Without Discount => R$ %.2f\n", s.IRWithoutDiscount)
	fmt.Printf("IR Discount         => R$ %.2f\n", s.IRDiscount)
	fmt.Printf("IR Value            => R$ %.2f\n", s.IR)
	fmt.Println("---------------------------------")
	fmt.Printf("Salary Net          => R$ %.2f\n", s.Net)
}
