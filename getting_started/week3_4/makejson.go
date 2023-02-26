package week3_4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func makeJson() {
	person := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello, There!")

	fmt.Print("Please, enter a name: ")
	scanner.Scan()
	person["name"] = scanner.Text()
	fmt.Printf("Entered name is: %s\n", person["name"])

	fmt.Print("Please, enter an address: ")
	scanner.Scan()
	person["address"] = scanner.Text()
	fmt.Printf("Entered address is: %s\n", person["address"])

	m, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// os.Stdout.Write(m)
	fmt.Printf("JSON format is: %s", m)
}
