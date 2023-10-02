package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m00n-arch/obnyashiwatel/internal/random"
	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	inputString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	swapLetters := random.RandomiseFunc(rules.SwapLetters, 0.80)
	inputDots := random.RandomiseFunc(rules.InputDots, 0.10)
	repeatLetters := random.RandomiseFunc(rules.RepeatLetters, 0.20)

	inputSliced := strings.Split(inputString, " ")
	for i := range inputSliced {
		inputSliced[i] = swapLetters(inputSliced[i])
		inputSliced[i] = inputDots(inputSliced[i])
		fmt.Print(repeatLetters(inputSliced[i]) + " ")

	}
}
