package main

import (
	"fmt"
	"sort"
)

func main() {
	angka := []int{1, 8, 5, 4, 3, 2, 6, 7}
	sort.Ints(angka)
	fmt.Printf("sorting: %d \n", angka)
}
