package rules

import (
	"math/rand"
)

const maxRepeats = 2

func RepeatLetters(inputText string) string {
	if random() {
		// Добавляем 1 к количеству повторов, чтобы избежать добавления пустых символов.
		return repeat(inputText, rand.Intn(maxRepeats)+1)
	} else {
		return inputText
	}
}

func repeat(inputText string, repeats int) string {
	for i := 0; i < repeats; i++ {
		inputText = string([]rune(inputText)[0]) + "-" + inputText
	}
	return inputText
}

const randChance = 3

func random() bool {
	return rand.Int()%randChance == 0
}
