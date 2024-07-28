package compiler

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	program := `
		fun main: num {
			let x 1
			let y 2
			let z (x + y)
			print(x)
			print(y)
			print(z)
			ret 0
		}
	`

	tr := NewTokenReader(program)
	tokens := tr.Tokenize()
	println(len(tokens), "tokens")
	for _, t := range tokens {
		println(t.String())
	}
}
