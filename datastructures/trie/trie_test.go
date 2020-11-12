package trie

import "testing"

var strs = []string{"foo", "bar", "bbc", "t", "tea", "teacap", "ten", "ten"}

func TestTrie(t *testing.T) {
	trie := New()
	for _, str := range strs {
		trie.Insert(str)
	}
	for _, str := range strs {
		count := trie.Search(str)
		if count == 0 {
			t.Errorf("want has %v but not", str)
			t.Errorf("get count %v", count)
		}
	}
	if count := trie.Search("ten"); count != 2 {
		t.Errorf("wanted count: %v", 2)
		t.Errorf("getted count: %v", count)
	}
}
