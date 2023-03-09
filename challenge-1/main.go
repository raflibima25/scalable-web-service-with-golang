package main

import (
	"fmt"
)

func main() {
	i := 21
	var j = true

	fmt.Println(i)
	fmt.Printf("%T \n", i)
	fmt.Println("%")

	fmt.Println(j)
	j = false
	fmt.Println(j)

	// fmt.Printf("%c \n", "я")
	// r, size := utf8.DecodeRuneInString("я")
	// 	fmt.Printf("%c %v\n", r, size)

	// rune, _ := utf8.DecodeRuneInString("я")
	// fmt.Println(rune)

	fmt.Printf("%U \n", []rune("Я"))

	fmt.Printf("%d \n", 21)
	fmt.Printf("%o \n", 25)
	fmt.Printf("%x \n", "f")
	fmt.Printf("%X \n", "F")

	// fmt.Printf("%U", "U+042F")
	fmt.Printf("%f \n", 123.456000)
	fmt.Printf("%e \n", 1.234560e+02)

}
