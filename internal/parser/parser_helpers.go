package parser

func (p *ProtoParser) extractKeyword() string {
	word := make([]rune, 0, 30)

	for !p.EOF() {
		if p.TestKeyword() {
			word = append(word, p.Next())
		} else {
			break
		}
	}

	return string(word)
}

func (p *ProtoParser) extractQuotedString() string {
	word := make([]rune, 0, 30)

	p.Peek('"') // TODO: we must check if error here
	for !p.EOF() {
		if p.Peek('"') {
			break
		} else {
			word = append(word, p.Next())
		}
	}

	return string(word)
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

func (p *ProtoParser) skipUntilNextLine() {
	for !p.EOF() {
		if !p.Peek('\n') {
			p.Next()
		} else {
			break
		}
	}
}
