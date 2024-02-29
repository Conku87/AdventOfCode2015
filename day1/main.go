package main

import (
	"fmt"
	"os"
)

const inputFile = "input.txt"

func main() {
	input := getInput()
	i := walkInput(input)
	fmt.Println(i)
}

func getInput() string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	} else {
		return string(content)
	}
}

func walkInput(input string) int {
	total := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			total += 1
		} else if input[i] == ')' {
			total -= 1
		}
	}
	return total
}
