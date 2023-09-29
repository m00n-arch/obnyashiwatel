package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	x, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(rules.SwapLetters(x))
}
