package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	inputString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	inputSliced := strings.Split(inputString, " ")
	for i := range inputSliced {
		inputSliced[i] = rules.SwapLetters(inputSliced[i])
		fmt.Print(rules.RepeatLetters(inputSliced[i]) + " ")

	}
}
