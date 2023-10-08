package domain

import (
	"strings"

	"github.com/m00n-arch/obnyashiwatel/internal/random"
	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func RunObnyash(inputText string) string {

	swapLetters := random.RandomiseFunc(rules.SwapLetters, 0.70)
	inputDots := random.RandomiseFunc(rules.AddDots, 0.10)
	repeatLetters := random.RandomiseFunc(rules.RepeatLetters, 0.20)
	insertSymbols := random.RandomiseFunc(rules.InsertSymbols, 0.20)

	textRunes := []rune(inputText)
	for j := range textRunes {
		k := string(textRunes[j])
		textRunes[j] = []rune(swapLetters(k))[0]
	}
	inputText = string(textRunes)

	res := ""
	inputWords := strings.Split(inputText, " ")
	for i := range inputWords {
		inputWords[i] = inputDots(inputWords[i])
		inputWords[i] = insertSymbols(inputWords[i])
		res = res + repeatLetters(inputWords[i]) + " "

	}
	return res
}
