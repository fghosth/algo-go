package algo_test

import (
	"fmt"
	"testing"

	"jvole.com/algo-go/trie"
)

func init() {

}

func TestTrie(test *testing.T) {
	t := trie.New()
	t.Add("foobare", 3)
	node, _ := t.Find("foobare")
	meta := node.Meta()
	fmt.Println(meta.(int))
	t.Add("derek", 33)
	t.Add("fookey", "sss")
	fmt.Println(t.PrefixSearch("foo"))
	fmt.Println(t.HasKeysWithPrefix("foo"))
	fmt.Println(t.FuzzySearch("fe"))

}
