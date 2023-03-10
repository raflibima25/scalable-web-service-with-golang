package main

import (
	"fmt"
)

func main() {

	// menampilkan nilai i = 21
	i := 21
	fmt.Println(i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	// menampilkan tanda %
	fmt.Println("%")

	// menampilkan nilai boolean j : true
	j := true
	fmt.Println(j)

	// menampilkan nilai boolean j : false
	j = false
	fmt.Println(j)

	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%U \n", []rune("Я"))

	// menampilkan nilai base 10 : 21
	fmt.Printf("%d \n", 21)

	// menampilkan nilai base 8 :25
	fmt.Printf("%o \n", 25)

	// menampilkan nilai base 16 : f
	fmt.Printf("%x \n", "f")

	// menampilkan nilai base 16 : F 13
	fmt.Printf("%X \n", "F")

	// menampilkan unicode karakter Я : U+042F
	fmt.Println("\u042F")

	// menampilkan float : 123.456000
	fmt.Printf("%f \n", 123.456000)

	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%e \n", 1.234560e+02)

}
