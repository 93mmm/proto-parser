package lexer

import (
	"github.com/93mmm/proto-parser/internal/errors"
	"github.com/93mmm/proto-parser/internal/parser/base"
	"github.com/93mmm/proto-parser/internal/parser/constants"
	"github.com/93mmm/proto-parser/internal/parser/source"
)

type Lexer struct {
	*base.BaseParser
}

func NewLexer(bp *base.BaseParser) *Lexer {
	return &Lexer{
		BaseParser: bp,
	}
}

func newTestLexer(input string) *Lexer {
	src := source.NewStringSource(input)
	bp := base.NewBaseParser(src)
	return NewLexer(bp)
}

func (p *Lexer) ExtractKeyword() (string, error) {
	p.SkipWhiteSpaces()

	keyword := make([]rune, 0, 30)
	for !p.EOF() {
		if p.TestKeyword() {
			keyword = append(keyword, p.Next())
		} else {
			break
		}
	}
	if len(keyword) == 0 {
		return "", errors.NewError(p.LineNumber(), p.CharNumber(), "Expected keyword, found %c", p.CurrentChar())
	}
	if !p.TestWhiteSpace() && !p.EOF() {
		return "", errors.NewError(p.LineNumber(), p.CharNumber(), "Expected end of keyword, found %c", p.CurrentChar())
	}
	return string(keyword), nil
}

func (p *Lexer) ExtractName() (string, error) {
	name := make([]rune, 0, 30)
	for !p.EOF() {
		if p.TestName() {
			name = append(name, p.Next())
		} else {
			break
		}
	}
	if len(name) == 0 {
		return "", errors.NewError(p.LineNumber(), p.CharNumber(), "Expected name, found %c", p.CurrentChar())
	}
	return string(name), nil
}

func (p *Lexer) ExtractQuotedString() (string, error) {
	if !p.Peek(constants.Quote) {
		return "", errors.NewError(p.LineNumber(), p.CharNumber(), "Quote expected, found %c", p.CurrentChar())
	}

	word := make([]rune, 0, 30)
	for !p.EOF() {
		switch {
		case p.Peek(constants.Quote):
			return string(word), nil
		case p.Test(constants.NextLine):
			return "", errors.NewError(p.LineNumber(), p.CharNumber(), "Not found end of quoted string, found \\n")
		default:
			word = append(word, p.Next())
		}
	}
	return "", errors.NewError(p.LineNumber(), p.CharNumber(), "EOF reached, nothing to extract")
}

// TODO: replace by skipping, NOT extraction
func (p *Lexer) ExtractNameBetweenParentheses() (string, error) {
	p.SkipWhiteSpaces()
	if err := p.PeekSymbol(constants.LeftParen); err != nil {
		return "", err
	}
	p.SkipWhiteSpaces()
	name, err := p.ExtractName()
	if err != nil {
		return "", err
	}
	p.SkipWhiteSpaces()
	if err := p.PeekSymbol(constants.RightParen); err != nil {
		return "", err
	}

	return name, nil
}

func (p *Lexer) SkipWhiteSpaces() {
	for !p.EOF() {
		if p.TestWhiteSpace() {
			p.Next()
		} else {
			break
		}
	}
}

func (p *Lexer) PeekSymbol(symbol rune) error {
	p.SkipWhiteSpaces()
	if !p.Peek(symbol) {
		return errors.NewError(p.LineNumber(), p.CharNumber(), "Expected %c found nothing", symbol)
	}
	return nil
}

func (p *Lexer) SkipUntilMatch(symbol rune) bool {
	for !p.EOF() {
		if !p.Test(symbol) {
			p.Next()
		} else {
			break
		}
	}
	if p.EOF() {
		return false
	}
	return true
}

// FIXME: bug, if eof reached and openCounter != 0 we don't throw error/return false
func (p *Lexer) SkipCurlyBraces() bool {
	if !p.SkipUntilMatch(constants.LeftBrace) {
		return false
	}
	p.Next()
	openCounter := 1
	for !p.EOF() && openCounter != 0 {
		switch {
		case p.Test(constants.RightBrace):
			p.Next()
			openCounter--
		case p.Test(constants.LeftBrace):
			p.Next()
			openCounter++
		default:
			p.Next()
		}
	}
	if openCounter != 0 {
		return false
	}
	p.Next()
	return true
}
