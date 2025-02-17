package t00139

func wordBreak(s string, wordDict []string) bool {
	return newTrie(wordDict).matches(s, nil)
}

type trie struct {
	children map[byte]*node
}

func newTrie(wordDict []string) *trie {
	t := &trie{
		children: make(map[byte]*node),
	}

	for _, w := range wordDict {
		t.addWord(w)
	}

	return t
}

func (t *trie) addWord(word string) {
	var n *node

	lastIndex := len(word) - 1

	for i, ch := range []byte(word) {
		if n == nil {
			var ok bool
			n, ok = t.children[ch]

			if !ok {
				n = newNode(ch)
				t.children[ch] = n
			}
		} else {
			n = n.getNodeOrCreate(ch)
		}

		n.setFinal(i == lastIndex)
	}
}

func (t *trie) matches(s string, cNode *node) bool {
	ch := s[0]

	var ok bool

	if cNode == nil {
		cNode, ok = t.children[ch]
	} else {
		cNode, ok = cNode.children[ch]
	}

	switch {
	case !ok:
		return false
	case len(s) == 1:
		return cNode.final
	case !cNode.final:
		return t.matches(s[1:], cNode)
	case cNode.final:
		return t.matches(s[1:], nil)
	default:
		panic("broken invariant")
	}
}

type node struct {
	value    byte
	final    bool
	children map[byte]*node
}

func newNode(
	value byte,
) *node {
	return &node{
		value:    value,
		final:    false,
		children: make(map[byte]*node),
	}
}

func (n *node) setFinal(final bool) {
	if n.final {
		return
	}

	n.final = final
}

func (n *node) getNodeOrCreate(value byte) *node {
	child, ok := n.children[value]

	if !ok {
		child = newNode(value)
		n.children[value] = child
	}

	return child
}
