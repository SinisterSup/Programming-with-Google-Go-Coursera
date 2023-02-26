package week2

import (
  "fmt"
)

func Trunc() {
  var num float64
  fmt.Println("Enter a float value to be truncated ")
  fmt.Scan(&num)
  fmt.Println("The truncated value is ", int32(num))
}
