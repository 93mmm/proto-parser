package parser

import "fmt"

// TODO: DRY principle, maybe fix it?

func (p *protoParser) extractKeyword() (string, error) {
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

func (p *protoParser) extractName() (string, error) {
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

func (p *protoParser) extractQuotedString() (string, error) {
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

func (p *protoParser) extractNameBetweenParentheses() (string, error) {
	p.skipWhiteSpaces()
	if err := p.peekOpenParenthesis(); err != nil {
		return "", err
	}
	p.skipWhiteSpaces()
	name, err := p.extractName()
	if err != nil {
		return "", err
	}
	p.skipWhiteSpaces()
	if err := p.peekCloseParenthesis(); err != nil {
		return "", err
	}

	return name, nil
}

func (p *protoParser) skipWhiteSpaces() {
	for !p.EOF() {
		if p.TestWhiteSpace() {
			p.Next()
		} else {
			break
		}
	}
}

func (p *protoParser) peekSymbol(symbol rune) error {
	p.skipWhiteSpaces()
	if !p.Peek(symbol) {
		return NewParserError(fmt.Sprintf("Expected %c found nothing", symbol), p.LineNumber(), p.CharNumber())
	}
	return nil
}

func (p *protoParser) peekSemicolon() error {
	return p.peekSymbol(';')
}

func (p *protoParser) peekEquals() error {
	return p.peekSymbol('=')
}

func (p *protoParser) peekOpenBrace() error {
	return p.peekSymbol('{')
}

func (p *protoParser) peekCloseBrace() error {
	return p.peekSymbol('}')
}

func (p *protoParser) peekOpenParenthesis() error {
	return p.peekSymbol('(')
}

func (p *protoParser) peekCloseParenthesis() error {
	return p.peekSymbol(')')
}

func (p *protoParser) skipUntilMatch(symbol rune) {
	for !p.EOF() {
		if !p.Test(symbol) {
			p.Next()
		} else {
			break
		}
	}
}
