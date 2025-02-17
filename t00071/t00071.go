package t00071

func simplifyPath(path string) string {
	tok := newTokenizer(path)
	return tok.simplify()
}

type tokenizer struct {
	pos  int
	path string
}

func newTokenizer(path string) *tokenizer {
	return &tokenizer{
		pos:  0,
		path: path,
	}
}

func (t *tokenizer) getNext() string {
	switch t.path[t.pos] {
	case '/':
		t.consumeChars('/')
		return "/"
	default:
		return t.consumeName()
	}
}

func (t *tokenizer) isActive() bool {
	return t.pos < len(t.path)
}

func (t *tokenizer) consumeChars(ch byte) {
	for t.isActive() && t.path[t.pos] == ch {
		t.pos++
	}
}

func (t *tokenizer) consumeName() string {
	fileName := ""

loop:
	for t.isActive() {
		next := t.path[t.pos]

		switch next {
		case '/':
			break loop
		default:
			fileName += string([]byte{next})
			t.pos++
		}
	}

	return fileName
}

func (t *tokenizer) simplify() string {
	tokens := make([]string, 0)

	for t.isActive() {
		token := t.getNext()

		switch {
		case token == "..":
			tokens = tokens[:max(len(tokens)-2, 1)]
		case token == ".":
			continue
		case token == "/" && len(tokens) > 0 && tokens[len(tokens)-1] == "/":
			continue
		case token == "/" && len(tokens) == 0:
			tokens = append(tokens, token)
		default:
			tokens = append(tokens, token)
		}
	}

	// trim last slash
	for len(tokens) > 1 && tokens[len(tokens)-1] == "/" {
		tokens = tokens[:len(tokens)-1]
	}

	res := ""

	for _, token := range tokens {
		res += token
	}

	return res
}
