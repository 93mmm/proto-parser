package parser

func (p *ProtoParser) extractKeyword() (string, error) {
	p.skipWhiteSpaces()

	word := make([]rune, 0, 30)
	for !p.EOF() {
		if p.TestKeyword() {
			word = append(word, p.Next())
		} else {
			break
		}
	}
	if len(word) == 0 {
		return "", NewParserError("Expected keyword, found nothing", p.LineNumber(), p.CharNumber())
	}
	return string(word), nil
}

func (p *ProtoParser) extractQuotedString() (string, error) {
	p.skipWhiteSpaces()

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
