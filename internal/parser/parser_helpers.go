package parser

import "fmt"

// TODO: DRY principle, maybe fix it?

func (p *ProtoParser) extractKeyword() (string, error) {
	p.skipWhiteSpaces()

	keyword := make([]rune, 0, 30)
	for !p.EOF() {
		if p.TestKeyword() {
			keyword = append(keyword, p.Next())
		} else {
			break
		}
	}
	if len(keyword) == 0 {
		return "", NewParserError("Expected keyword, found nothing", p.LineNumber(), p.CharNumber())
	}
	return string(keyword), nil
}

func (p *ProtoParser) extractName() (string, error) {
	name := make([]rune, 0, 30)
	for !p.EOF() {
		if p.TestName() {
			name = append(name, p.Next())
		} else {
			break
		}
	}
	if len(name) == 0 {
		return "", NewParserError("Expected name, found nothing", p.LineNumber(), p.CharNumber())
	}
	return string(name), nil
}

func (p *ProtoParser) extractQuotedString() (string, error) {
	if !p.Peek('"') {
		return "", NewParserError("Quote expected", p.LineNumber(), p.CharNumber())
	}

	word := make([]rune, 0, 30)
	for !p.EOF() {
		switch {
		case p.Peek('"'):
			return string(word), nil
		case p.Test('\n'):
			return "", NewParserError("Not found end of quoted string", p.LineNumber(), p.CharNumber())
		default:
			word = append(word, p.Next())
		}
	}
	return "", NewParserError("EOF reached, nothing to extract", p.LineNumber(), p.CharNumber())
}

func (p *ProtoParser) skipWhiteSpaces() {
	for !p.EOF() {
		if p.TestWhiteSpace() {
			p.Next()
		} else {
			break
		}
	}
}

func (p *ProtoParser) peekSymbol(symbol rune) error {
	p.skipWhiteSpaces()
	if !p.Peek(symbol) {
		return NewParserError(fmt.Sprintf("Expected %c found nothing", symbol), p.LineNumber(), p.CharNumber())
	}
	return nil
}

func (p *ProtoParser) peekSemicolon() error {
	return p.peekSymbol(';')
}

func (p *ProtoParser) peekEquals() error {
	return p.peekSymbol('=')
}

func (p *ProtoParser) peekOpenCurlyBrace() error {
	return p.peekSymbol('{')
}

func (p *ProtoParser) peekCloseCurlyBrace() error {
	return p.peekSymbol('}')
}

// TODO: maybe we don't need it?
func (p *ProtoParser) skipUntilNextLine() {
	for !p.EOF() {
		if !p.Peek('\n') {
			p.Next()
		} else {
			break
		}
	}
}
