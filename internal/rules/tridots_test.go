package rules_test

import (
	"testing"

	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func TestAddDots(t *testing.T) {
	input := "больше"
	exp := "больше..."
	out := rules.AddDots(input)
	if out != exp {
		t.Fatalf("expected: %s, got: %s", exp, out)
	}
}
