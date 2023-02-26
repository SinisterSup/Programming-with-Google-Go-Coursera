package course2

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
) 

type Animal1 interface {
  Eat()
  Move()
  Speak()
}

type AnimalAttr struct {
  isWhat string
  food string
  locomotion string
  sound string
}

func (a *AnimalAttr) Eat() {
  fmt.Println("Your animal eats -", a.food)
}

func (a *AnimalAttr) Move() {
  fmt.Println("The locomotion method of your Animal is -", a.locomotion)
}

func (a *AnimalAttr) Speak() {
  fmt.Println("Your Animal makes the sound -", a.sound)
}

// func (a *AnimalAttr) IsWhat() {
//   fmt.Println("Your Animal is a ", a.isWhat)
// }

// func main() {
//   animalNames := make(map[string]Animal1)
//   for {
//     fmt.Println("> ")
//     reader := bufio.NewReader(os.Stdin)
//     input, err := reader.ReadString('\n')
//     if err != nil {
//       panic(err)
//     }
//     input = strings.TrimSpace(input) // remove the newline character
//     if input == "X" || input == "x" { break }
//     // fmt.Println("The input you provided is -", input)
//     info := strings.Split(input, " ")


//     if info[0] == "newanimal" {
//       createNewAnimal(info[1], info[2], animalNames)
//     } else if info[0] == "query" {
//       performAction(info[1], info[2], animalNames)
//     } else {
//       fmt.Println("Please Enter a valid first command among (newanimal or query)")
//     }
//   
//   }
// }

func createNewAnimal(name, anim string, animalNames map[string]Animal1) {
  var newAnimal Animal1

  switch anim {
  case "cow":
    newAnimal = Animal1(&AnimalAttr{
		    isWhat:     "cow",
		    food:       "grass",
		    locomotion: "walk",
		    sound:      "moo",
	  })
    fmt.Println("Created It!")

  case "bird":
    newAnimal = Animal1(&AnimalAttr{
		    isWhat:     "bird",
		    food:       "worms",
		    locomotion: "fly",
		    sound:      "peep",
	  })
    fmt.Println("Created It!")
  case "snake":
    newAnimal = Animal1(&AnimalAttr{
		    isWhat:     "snake",
		    food:       "mice",
		    locomotion: "slither",
		    sound:      "hsss",
	  })
    fmt.Println("Created It!")
  default:
    fmt.Println("Enter a valid animal type among (cow, bird, snake)")
  }

  val, ok := animalNames[name]
  if !ok {
    animalNames[name] = newAnimal
    fmt.Println("The Animals now present are :", animalNames)
  } else {
    fmt.Printf("The animal with the name %s is already there! and it is a %v\n", name, val)
  }
}

func performAction(name, action string, animalNames map[string]Animal1) {
  fmt.Println(" The animals already there are -", animalNames) 
  myAnimal, ok := animalNames[name]
  if ok {
    switch action {
    case "eat":
      myAnimal.Eat()
    case "move":
      myAnimal.Move()
    case "speak":
      myAnimal.Speak()
    default:
      fmt.Println("Enter a valid command among (eat, move, speak) for the animal to act")
    }
  } else {
    fmt.Println("The Animal with this names has not yet been created !")
  }
}
