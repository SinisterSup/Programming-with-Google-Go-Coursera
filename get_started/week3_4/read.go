package week3_4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PersonName struct {
	fname string
	lname string
}

func readFile() {
	var fileName string

	fmt.Print("Enter the file name (with .txt): ")
	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	var personNames []PersonName
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		words := strings.Split(line, " ")
		if len(words) == 2 {
			firstName, lastName := words[0], words[1]
			personNames = append(personNames, PersonName{fname: firstName, lname: lastName})
			// fmt.Printf("First Name: %s, Last Name: %s \n", firstName, lastName)
		} else {
			fmt.Println("Invalid line:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	for _, name := range personNames {
		fmt.Printf("First Name: %s, Last Name: %s \n", name.fname, name.lname)
	}

}