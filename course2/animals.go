package course2

import (
  "fmt"
  // "bufio"
  // "os"
  "strings"
)

type Animal struct {
  food string
  locomotion string
  noise string
}

func (a *Animal) createAnimal(food, motion, sound string) {
  a.food = food
  a.locomotion = motion
  a.noise = sound
}

func (a *Animal) Eat() {
  fmt.Println("Your Animal eats -", a.food)
}

func (a *Animal) Move() {
  fmt.Println("The locomotion method of your Animal is -", a.locomotion)
}

func (a *Animal) Speak() {
  fmt.Println("Your Animal makes the sound -", a.noise)
}

// func main() {
//   for {
//     fmt.Println("> ")
//     reader := bufio.NewReader(os.Stdin)
//     input, err := reader.ReadString('\n')
//     if err != nil {
//       panic(err)
//     }
//     input = strings.TrimSpace(input) // remove the newline character

//     if input == "X" || input == "x" { break }

//     fmt.Println("The input you provided is -", input)
//   
//     animal_action := strings.Split(input, " ")
//     animal, action := animal_action[0], animal_action[1]

//     performAction(animal , action)
//   
//   }
// }

func performAction(animal, command string) {
  animal = strings.ToLower(animal)
  command = strings.ToLower(command)
  
  var ani Animal
  
  switch animal {
    case "cow": 
      ani.createAnimal("grass", "walk", "moo")
    case "bird":
      ani.createAnimal("worms", "fly", "peep")
    case "snake":
      ani.createAnimal("mice", "slither", "hsss")
    default:
      fmt.Printf("The animal %s you entered is not in the list \n", animal)
  }

  switch command {
    case "eat":
      ani.Eat()
    case "move":
      ani.Move()
    case "speak":
      ani.Speak()
    default:
      fmt.Println("Please enter a valid command among (eat, move, speak)")
  }
}
