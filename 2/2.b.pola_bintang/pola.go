package main

import "fmt"

func main() {
	looping := 1
	for i := 1; i <= 5; i++ {
		fmt.Println("")
		for j := 1; j <= looping; j++ {
			fmt.Print("*")
		}
		looping++
	}
}
