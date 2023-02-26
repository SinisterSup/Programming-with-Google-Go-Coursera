package week3_4

import (
	"fmt"
	"sort"
)

func SortSlice() {
	var nums []int
	var inputNum int
	// var n int
	var inputChar string
	
	// the program shall quit only if the user enters a 'X' or 'x' instead of an integer
	for {
		fmt.Println("Enter an integer (X to quit): ")
		fmt.Scan(&inputChar)
		if inputChar == "X" || inputChar == "x" {
			break
		}
		fmt.Sscan(inputChar, &inputNum)
		nums = append(nums, inputNum)
		sort.Ints(nums)
		fmt.Println("Sorted numbers: ", nums)
	}
	// for i := 0; i < n; i++ {
	// 	fmt.Print("Enter the number: ")
	// 	fmt.Scan(&inputNum)
	// 	nums = append(nums, inputNum)
	// 	sort.Ints(nums)
	// 	fmt.Println("Sorted numbers: ", nums)
	// }
}