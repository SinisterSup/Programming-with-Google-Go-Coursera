package course2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "strings"
)

// func main() {
//   array := arrInput()
//   BubbleSort(array)
//   fmt.Println("The sorted array is", array)
// }

func BubbleSort(nums []int) []int {
  n := len(nums)
  for i := 0; i < n; i++ {
    swapped := false
    for j := 0; j < n - 1 - i; j++ {
      if nums[j] > nums[j+1] {
        Swap(nums, j)
        swapped = true
      }
    }
    if !swapped {
      break
    }
  }
  return nums
}

func Swap(sli []int, index int) {
  sli[index], sli[index+1] = sli[index+1], sli[index]
}

func arrInput() []int {
  fmt.Println("Please Enter space separated Integers ")

  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
  if err != nil {
    panic(err)
  }
  input = strings.TrimSpace(input) // remove the newline character
  fmt.Println("The input you provided is -", input)
  
  strNums := strings.Split(input, " ")

  fmt.Println(strNums)

  var numbers []int

  for _, strNum := range strNums {
    num, err := strconv.Atoi(strNum)
    if err != nil {
      panic(err)
    }
    numbers = append(numbers, num)
  }
  
  return numbers

}
