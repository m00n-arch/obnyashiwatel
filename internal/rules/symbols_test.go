package rules_test

import (
	"testing"

	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func TestInsertSymbols(t *testing.T) {
	input := "больше"
	exp := "больше *UWU*"

	out := rules.InsertSymbols(input)
	if out != exp {
		t.Fatalf("expected: %s, got: %s", exp, out)
	}
}
