package course2

import (
  // "fmt"
  "math"
)

// func main() {
//   var a, vo, so float64
//   var time float64

//   fmt.Println("Please Enter the acceleration, initial velocity, and initial displacment: ")
//   _, err := fmt.Scanln(&a, &vo, &so)
//   if err != nil {
//     fmt.Println("Enter only float values ", err)
//   }

//   fmt.Println("Now Enter the time to compute Displacement after time t -")
//   _, e := fmt.Scanln(&time)
//   if e != nil {
//     fmt.Println("ENTER float value for time ", err)
//   }

//   eqn := GenDisplaceFn(a, vo, so)

//   displacement := eqn(time)
//   fmt.Println("The Displacement after time t for the given acc, velocity, and initial displacment is :", displacement)

// }

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
  return func (t float64) float64 { 
    return (0.5 * a * math.Pow(t, 2)) + (v0 * t) + s0 
  }
}
