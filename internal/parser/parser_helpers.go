package parser

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
		return "", NewParserError(p.LineNumber(), p.CharNumber(), "Expected keyword, found %c", p.CurrentChar())
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
		return "", NewParserError(p.LineNumber(), p.CharNumber(), "Expected name, found %c", p.CurrentChar())
	}
	return string(name), nil
}

func (p *protoParser) extractQuotedString() (string, error) {
	if !p.Peek('"') {
		return "", NewParserError(p.LineNumber(), p.CharNumber(), "Quote expected, found %c", p.CurrentChar())
	}

	word := make([]rune, 0, 30)
	for !p.EOF() {
		switch {
		case p.Peek('"'):
			return string(word), nil
		case p.Test('\n'):
			return "", NewParserError(p.LineNumber(), p.CharNumber(), "Not found end of quoted string, found \\n")
		default:
			word = append(word, p.Next())
		}
	}
	return "", NewParserError(p.LineNumber(), p.CharNumber(), "EOF reached, nothing to extract")
}

func (p *protoParser) extractNameBetweenParentheses() (string, error) {
	p.skipWhiteSpaces()
	if err := p.peekSymbol('('); err != nil {
		return "", err
	}
	p.skipWhiteSpaces()
	name, err := p.extractName()
	if err != nil {
		return "", err
	}
	p.skipWhiteSpaces()
	if err := p.peekSymbol(')'); err != nil {
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
		return NewParserError(p.LineNumber(), p.CharNumber(), "Expected %c found nothing", symbol)
	}
	return nil
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

func (p *protoParser) skipCurlyBraces() {
	p.skipUntilMatch('{')
	p.Next()
	openCounter := 1
	for !p.EOF() && openCounter != 0 {
		switch {
		case p.Test('}'):
			p.Next()
			openCounter--
		case p.Test('{'):
			p.Next()
			openCounter++
		default:
			p.Next()
		}
	}
	p.Next()
}
