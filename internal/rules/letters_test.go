package rules_test

import (
	"testing"

	"github.com/m00n-arch/obnyashiwatel/internal/rules"
)

func TestSwapLetters(t *testing.T) {
	input := "больше всего на свете я Люблю кушать нефть, пить ртуть и закусывать все это булочкой с ураном-239"
	exp := "бовьфе ффего на ффете я Вюбвю куфать нефть, пить лтуть и закуфыфать ффе это бувочкой ф уланом-239"
	out := rules.SwapLetters(input)
	if out != exp {
		t.Fatalf("expected: %s, got: %s", exp, out)
	}
}
