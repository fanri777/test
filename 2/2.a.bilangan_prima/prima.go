package main

import "fmt"

func main() {
	for prima := 1; prima < 100; prima++ {
		i := 0
		for bag := 1; bag < 100; bag++ {
			if prima%bag == 0 {
				i++
			}
		}
		if (i == 2) && (prima != 1) {
			fmt.Println(prima)
		}
	}
}
