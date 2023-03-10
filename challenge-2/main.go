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
				fmt.Printf("character %U '%c' starts at byte position %d \n", v, v, i)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
