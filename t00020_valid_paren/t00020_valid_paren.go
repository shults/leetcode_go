package t00020_valid_paren

func isValid(s string) bool {
	st := newStack()

	for _, paren := range []byte(s) {
		switch {
		case paren == '{' || paren == '(' || paren == '[':
			st.push(paren)
		case paren == '}' || paren == ')' || paren == ']':
			if st.isEmpty() {
				return false
			}

			item := st.pop()

			if !areOpposite(item, paren) {
				return false
			}
		default:
			panic("unexpected paren")
		}
	}

	return st.isEmpty()
}

var oppositeCases = []struct{ a, b byte }{
	{'{', '}'},
	{'[', ']'},
	{'(', ')'},
}

func areOpposite(parenA, parenB byte) bool {
	for _, c := range oppositeCases {
		if c.a == parenA && c.b == parenB {
			return true
		}
	}

	return false
}

type stack struct {
	items []byte
}

func newStack() *stack {
	return &stack{
		items: make([]byte, 0),
	}
}

func (s *stack) push(item byte) {
	s.items = append(s.items, item)
}

func (s *stack) size() int {
	return len(s.items)
}

func (s *stack) isEmpty() bool {
	return s.size() == 0
}

func (s *stack) pop() byte {
	last := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return last
}
