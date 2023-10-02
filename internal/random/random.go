package random

import (
	"math/rand"
)

func RandomiseFunc(rule func(string) string, chance float64) func(string) string {
	return func(text string) string {
		if rand.Float64() < chance {
			return rule(text)
		} else {
			return text
		}
	}
}
