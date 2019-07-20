package main

import (
	"fmt"
	"sort"
)

func bar(maxB int, colsB int, numB []int) {
	for i := 0; i < maxB; i++ {
		for r := 0; r < colsB; r++ {
			var k = i + numB[r]
			if k >= maxB {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("\t")
			// fmt.Print(num[r])
		}
		fmt.Println()
	}
}

func main() {
	var num = []int{1, 4, 5, 6, 8, 2}
	var max = 8
	var cols = 6
	bar(max, cols, num)
	fmt.Println(num)
	// Ascending
	sort.Slice(num, func(i, v int) bool {
		return num[i] < num[v]
	})
	bar(max, cols, num)
	for _, each := range num {
		fmt.Printf("%d ", each)
	}
	fmt.Print("\n")
	// Descending
	sort.Slice(num, func(i, v int) bool {
		return num[i] > num[v]
	})
	bar(max, cols, num)
	for _, each := range num {
		fmt.Printf("%d ", each)
	}
	fmt.Print("\n")
}
