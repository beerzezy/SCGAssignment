package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/beerzezy/SCGTest/assignment"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var args string

	fmt.Print("Enter ARGS: ")

	if scanner.Scan() {
		args = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result, err := assignment.Calculation(args)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", result)
}
