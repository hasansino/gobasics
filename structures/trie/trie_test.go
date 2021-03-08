package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("helio")
	trie.AddWord("hello")
	trie.AddWord("hallo")
	trie.AddWord("hectic")
	trie.AddWord("haptic")

	assert.EqualValues(t, []string{"helio", "hello", "hectic"}, trie.FindPartial("he"))
	assert.EqualValues(t, []string{"hallo", "haptic"}, trie.FindPartial("ha"))
	assert.EqualValues(t, []string{"helio", "hello", "hectic", "hallo", "haptic"}, trie.FindPartial(""))
	assert.Nil(t, trie.FindPartial("foo"))
}
