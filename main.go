package main

import (
	"fmt"

	math "hello/color"

	cl "github.com/fatih/color"
)

func main() {
	var S int
	fmt.Print("введите натуральное чётное число (ввод нечётного= округление):")
	fmt.Scan(&S)
	R := cl.New(cl.FgRed, cl.Bold)
	G := cl.New(cl.FgGreen, cl.Bold)
	B := cl.New(cl.FgBlue, cl.Bold)
	Y := cl.New(cl.FgYellow, cl.Bold)
	cl.Red("разделил на 2:")
	R.Println(math.Pazdelitna2(S))
	cl.Green("умножил на 2:")
	G.Println(math.Ymnohitna2(S))
	cl.Blue("убавил 10:")
	B.Println(math.Minys10(S))
	cl.Yellow("прибавил 10:")
	Y.Println(math.Nlus10(S))
}
