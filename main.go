package main

import "fmt"

func main() {
	numbers()
	text()
	addnumber()
}

func numbers() {

	var A int
	A = 10
	var B float32
	B = 10.002

	fmt.Printf("%d\n", A)
	fmt.Println("Float:", B)

}
func text() {
	var C string
	fmt.Print("Enter a string:")
	fmt.Scan(&C)
	fmt.Printf("%s\n", C)
}
func addnumber() {
	var H, J, I int
	fmt.Print("Enter a number:\n")
	fmt.Scan(&H, &J)
	I = H + J
	fmt.Printf("Output:%d", I)

}
