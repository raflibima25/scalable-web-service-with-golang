package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			str := "CAPABLECANYOU"
			for i, v := range str {
				if i%2 == 1 {
					continue
				}
				fmt.Printf("character %#U starts at byte position %d \n", v, i)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}

	// another ver
	/*
		for j := 0; j <= 10: j++ {
			if j == 5 {
				teks := []rune('С','А','Ш','А','Р','В','О')
				for indeks := 0; indeks < len(teks); indeks++ {
					char := teks[indeks]
					fmt.Printf("character %U '%s' start at byte position %d \n", char, string(char), indeks*2)
				}
			} else {
				fmt.Println("Nilai j = ", j)
			}
		}

	*/
}
