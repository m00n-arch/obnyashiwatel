package rules_test

import (
	"regexp"
	"testing"

	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func TestRepeatLetters(t *testing.T) {
	input := "больше"
	exp := regexp.MustCompile("(б-)*больше")
	out := rules.RepeatLetters(input)
	if exp.MatchString(out) == false {
		t.Fatalf("expected: %s, got: %s", exp.String(), out)
	}
}
