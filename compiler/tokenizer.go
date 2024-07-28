package compiler

import (
	"bufio"
	"io"
	"strings"

	"github.com/dixtel/m/utils"
)

type TokenID int

const (
	KEYWORD_FUN TokenID = iota
	KEYWORD_LET TokenID = iota

	IDENTIFIER TokenID = iota
	NUMBER     TokenID = iota

	BEGIN_BRACKET     TokenID = iota
	END_BRACKET       TokenID = iota
	BEGIN_PARENTHESIS TokenID = iota
	END_PARENTHESIS   TokenID = iota

	COLON   TokenID = iota
	PLUS    TokenID = iota
	MINUS   TokenID = iota
	EOF     TokenID = iota
	UNKNOWN TokenID = iota
)

// Source: https://doc.rust-lang.org/reference/whitespace.html
var WHITE_SPACES = []rune{
	rune(0x0009), // U+0009 horizontal tab, '\t'
	rune(0x000A), // U+000A line feed, '\n'
	rune(0x000B), // U+000B vertical tab
	rune(0x000C), // U+000C form feed
	rune(0x000D), // U+000D carriage return, '\r'
	rune(0x0020), // U+0020 space, ' '
	rune(0x0085), // U+0085 next line
	rune(0x200E), // U+200E left-to-right mark
	rune(0x200F), // U+200F right-to-left mark
	rune(0x2028), // U+2028 line separator
	rune(0x2029), // U+2029 paragraph separator
}

var START_OF_IDENTIFIER = utils.Spread(
	generateRunes('a', 'z'),
	generateRunes('A', 'Z'),
	[]rune{rune('_'), rune('$')},
)

var REST_OF_IDENTIFIER = utils.Spread(
	generateRunes('a', 'z'),
	generateRunes('A', 'Z'),
	generateRunes('0', '9'),
	[]rune{rune('_')},
)

var BEGIN_NUMBER = generateRunes('0', '9')

var REST_OF_NUMBER = utils.Spread(
	generateRunes('0', '9'),
	[]rune{rune('_')},
)

type Token struct {
	Begin   int
	Text    string
	TokenID TokenID
}

func (t Token) String() string {
	switch t.TokenID {
	case KEYWORD_FUN:
		return "KEYWORD_FUN " + t.Text
	case KEYWORD_LET:
		return "KEYWORD_LET " + t.Text
	case IDENTIFIER:
		return "IDENTIFIER " + t.Text
	case NUMBER:
		return "NUMBER " + t.Text
	case BEGIN_BRACKET:
		return "BEGIN_BRACKET " + t.Text
	case END_BRACKET:
		return "END_BRACKET " + t.Text
	case BEGIN_PARENTHESIS:
		return "BEGIN_PARENTHESIS " + t.Text
	case END_PARENTHESIS:
		return "END_PARENTHESIS " + t.Text
	case COLON:
		return "COLON " + t.Text
	case PLUS:
		return "PLUS " + t.Text
	case MINUS:
		return "MINUS " + t.Text
	case EOF:
		return "EOF " + t.Text
	}

	return "unknown: " + t.Text
}

type TokenReader struct {
	currIdx int
	buff    *bufio.Reader
	tokens  []Token
}

func NewTokenReader(file string) *TokenReader {
	buff := bufio.NewReader(strings.NewReader(file))
	return &TokenReader{
		buff:    buff,
		currIdx: -1,
	}
}

func (tr *TokenReader) Tokenize() []Token {
	for {
		token := tr.nextToken()
		tr.tokens = append(tr.tokens, token)

		if token.TokenID == EOF {
			break
		}
	}

	return tr.tokens
}

func (tr *TokenReader) nextToken() Token {
	tr.skipWhiteSpaces()
	s, err := tr.readRune()
	if err == io.EOF {
		return Token{
			Begin:   tr.currIdx,
			Text:    "",
			TokenID: EOF,
		}
	}

	switch s {
	case '+':
		return Token{
			Begin:   tr.currIdx,
			TokenID: PLUS,
			Text:    string(s),
		}
	case '-':
		return Token{
			Begin:   tr.currIdx,
			TokenID: MINUS,
			Text:    string(s),
		}
	case ':':
		return Token{
			Begin:   tr.currIdx,
			TokenID: COLON,
			Text:    string(s),
		}
	case '{':
		return Token{
			Begin:   tr.currIdx,
			TokenID: BEGIN_BRACKET,
			Text:    string(s),
		}
	case '}':
		return Token{
			Begin:   tr.currIdx,
			TokenID: END_BRACKET,
			Text:    string(s),
		}
	case '(':
		return Token{
			Begin:   tr.currIdx,
			TokenID: BEGIN_PARENTHESIS,
			Text:    string(s),
		}
	case ')':
		return Token{
			Begin:   tr.currIdx,
			TokenID: END_PARENTHESIS,
			Text:    string(s),
		}
	}

	if utils.Has(BEGIN_NUMBER, s) {
		rest := tr.readUntilNot(REST_OF_NUMBER...)
		text := string(s) + rest
		return Token{
			Begin:   tr.currIdx,
			TokenID: NUMBER,
			Text:    text,
		}
	}

	if utils.Has(START_OF_IDENTIFIER, s) {
		rest := tr.readUntilNot(REST_OF_IDENTIFIER...)
		text := string(s) + rest

		if text == "fun" {
			return Token{
				Begin:   tr.currIdx,
				TokenID: KEYWORD_FUN,
				Text:    text,
			}
		}

		if text == "let" {
			return Token{
				Begin:   tr.currIdx,
				TokenID: KEYWORD_LET,
				Text:    text,
			}
		}

		return Token{
			Begin:   tr.currIdx,
			TokenID: IDENTIFIER,
			Text:    text,
		}
	}

	unknown := tr.readUntilNot(WHITE_SPACES...)

	tr.tokens = append(tr.tokens, Token{
		Begin:   tr.currIdx,
		Text:    unknown,
		TokenID: UNKNOWN,
	})

	return tr.nextToken()
}

func (tr *TokenReader) skipWhiteSpaces() error {
	r, err := tr.readRune()
	if err != nil {
		return err
	}

	if utils.Has(WHITE_SPACES, r) {
		return tr.skipWhiteSpaces()
	}

	err = tr.buff.UnreadRune()
	tr.currIdx -= 1
	if err != nil {
		panic("this cannot happen: " + err.Error())
	}

	return nil
}

func (tr *TokenReader) readRune() (rune, error) {
	r, _, err := tr.buff.ReadRune()
	tr.currIdx += 1
	if err != nil {
		return rune(0x00), err
	}

	return r, nil
}

func (tr *TokenReader) isEOF() bool {
	_, _, err := tr.buff.ReadRune()
	tr.currIdx += 1
	return err == io.EOF
}

func (tr *TokenReader) unreadRune() {
	err := tr.buff.UnreadRune()
	tr.currIdx -= 1
	if err != nil {
		panic("cannot unread utf8 character: " + err.Error())
	}
}

func (tr *TokenReader) readUntilNot(ch ...rune) (ret string) {
	for {
		r, _, err := tr.buff.ReadRune()
		if err != nil {
			panic("cannot read utf8 character: " + err.Error())
		}

		if !utils.Has(ch, r) {
			tr.unreadRune()
			return
		}

		tr.currIdx += 1
		ret += string(r)
	}
}

func generateRunes(beginIdx, endIdx int) (ret []rune) {
	for i := beginIdx; i <= endIdx; i++ {
		ret = append(ret, rune(i))
	}
	return
}
