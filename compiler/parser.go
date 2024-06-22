package compiler

import "github.com/dixtel/m/utils"

type Token int

const (
	KEYWORD_FUN       Token = iota
	KEYWORD_LET       Token = iota
	RIGHT_ARROW       Token = iota
	START_BRACKET     Token = iota
	END_BRACKET       Token = iota
	START_PARENTHESIS Token = iota
	END_PARENTHESIS   Token = iota
	IDENTIFIER        Token = iota
	PLUS              Token = iota
	MINUS             Token = iota
	DOLLAR_SIGN       Token = iota
)

// Source: https://doc.rust-lang.org/reference/whitespace.html
var WHITE_SPACES = []rune{
	rune(0x0009), // U+0009 (horizontal tab, '\t')
	rune(0x000A), // U+000A (line feed, '\n')
	rune(0x000B), // U+000B (vertical tab)
	rune(0x000C), // U+000C (form feed)
	rune(0x000D), // U+000D (carriage return, '\r')
	rune(0x0020), // U+0020 (space, ' ')
	rune(0x0085), // U+0085 (next line)
	rune(0x200E), // U+200E (left-to-right mark)
	rune(0x200F), // U+200F (right-to-left mark)
	rune(0x2028), // U+2028 (line separator)
	rune(0x2029), // U+2029 (paragraph separator)
}

func removeWhiteSpaces(in string) (ret string) {
	for _, r := range in {
		if utils.Has(WHITE_SPACES, r) {
			continue
		}
	}

	return ""
}

func Parse(input string) []Token {

}


type TokenReader struct {
	src string
	size int
	

}