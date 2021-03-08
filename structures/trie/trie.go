package trie

// https://en.wikipedia.org/wiki/Trie#:~:text=In%20computer%20science%2C%20a%20trie,key%2C%20but%20by%20individual%20characters.

// Trie is prefix tree
type Trie struct {
	root *Node
}

// Node of trie
type Node struct {
	final         bool
	value         string
	parent        *Node
	children      map[rune]*Node
	childrenList  []*Node
	finalChildren int32
}

func (n *Node) incrementFinals() {
	for p := n.parent; p != nil; p = p.parent {
		p.finalChildren++
	}
}

// NewTrie creates empty trie with initialized root node
func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[rune]*Node),
		},
	}
}

// AddWord to trie
func (t *Trie) AddWord(w string) {
	t.addWord(t.root, []rune(w), "")
}

func (t *Trie) addWord(n *Node, w []rune, prefix string) {
	if n == nil {
		return
	}

	if len(w) == 0 && len(prefix) > 0 {
		n.final = true
		n.incrementFinals()
		return
	}

	prefix = n.value

	if _, ok := n.children[w[0]]; ok {
		t.addWord(n.children[w[0]], w[1:], prefix)
	} else {
		newNode := &Node{
			parent:   n,
			value:    prefix + string(w[0]),
			children: make(map[rune]*Node),
		}

		n.children[w[0]] = newNode
		n.childrenList = append(n.childrenList, newNode)

		if len(w[1:]) > 0 {
			t.addWord(n.children[w[0]], w[1:], prefix)
		} else {
			newNode.final = true
			newNode.incrementFinals()
		}
	}
}

// FindPartial searches for strings that matches (pat.*)
func (t *Trie) FindPartial(pat string) []string {
	if t.root == nil {
		return nil
	}

	var (
		runes    = []rune(pat)
		currNode = t.root
	)

	for j := 0; j < len(runes); j++ {
		r := runes[j]
		if _, ok := currNode.children[r]; ok {
			currNode = currNode.children[r]
		} else {
			return nil
		}
	}

	return t.findPartial(currNode)
}

func (t *Trie) findPartial(n *Node) []string {
	if len(n.children) == 0 {
		return []string{n.value}
	}
	var matches = make([]string, 0)
	if n.final {
		matches = append(matches, n.value)
	}
	for j := 0; j < len(n.childrenList); j++ {
		matches = append(matches, t.findPartial(n.childrenList[j])...)
	}
	return matches
}
