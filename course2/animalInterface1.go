package course2

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
) 

type Animal interface {
  Eat()
  Move()
  Speak()
}

type Cow struct {
  food string
  locomotion string
  sound string
}
type Bird struct {
  food string
  locomotion string
  sound string
}
type Snake struct {
  food string
  locomotion string
  sound string
}

func (c *Cow) Eat() {
  fmt.Println("Your animal eats -", c.food)
}
func (c *Cow) Move() {
  fmt.Println("The locomotion method of your Animal is -", c.locomotion)
}
func (c *Cow) Speak() {
  fmt.Println("Your Animal makes the sound -", c.sound)
}

func (b *Bird) Eat() {
  fmt.Println("Your animal eats -", b.food)
}
func (b *Bird) Move() {
  fmt.Println("The locomotion method of your Animal is -", b.locomotion)
}
func (b *Bird) Speak() {
  fmt.Println("Your Animal makes the sound -", b.sound)
}

func (s *Snake) Eat() {
  fmt.Println("Your animal eats -", s.food)
}
func (s *Snake) Move() {
  fmt.Println("The locomotion method of your Animal is -", s.locomotion)
}
func (s *Snake) Speak() {
  fmt.Println("Your Animal makes the sound -", s.sound)
}
// func (a *AnimalAttr) IsWhat() {
//   fmt.Println("Your Animal is a ", a.isWhat)
// }

// func main() {
//   animalNames := make(map[string]Animal)
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
//       createAnimal(info[1], info[2], animalNames)
//     } else if info[0] == "query" {
//       queryAnimal(info[1], info[2], animalNames)
//     } else {
//       fmt.Println("Please Enter a valid first command among (newanimal or query)")
//     }
//   
//   }
// }

func createAnimal(name, anim string, animalNames map[string]Animal) {
  var newAnimal Animal

  switch anim {
  case "cow":
    newAnimal = Animal(&Cow{
		    food:       "grass",
		    locomotion: "walk",
		    sound:      "moo",
	  })
    fmt.Println("Created It!")

  case "bird":
    newAnimal = Animal(&Bird{
		    food:       "worms",
		    locomotion: "fly",
		    sound:      "peep",
	  })
    fmt.Println("Created It!")
  case "snake":
    newAnimal = Animal(&Snake{
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

func queryAnimal(name, action string, animalNames map[string]Animal) {
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
