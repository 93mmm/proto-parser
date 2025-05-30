package parser

func (p *ProtoParser) extractWord() string {
	word := make([]rune, 0, 30)

	for !p.EOF() {
		if p.PeekSymbol() {
			word = append(word, p.Next())
		} else {
			break
		}
	}
	return string(word)
}

func (p *ProtoParser) skipWhiteSpaces() {
	for !p.EOF() {
		if p.PeekWhiteSpace() {
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
