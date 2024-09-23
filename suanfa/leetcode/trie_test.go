package main

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if got := trie.Search("apple"); got != true {
		t.Errorf("Search() = %v, want %v", got, true)
	}

	s2 := trie.Search("app") // 返回 False
	t.Logf("search app %v", s2)

	s3 := trie.StartsWith("app") // 返回 True
	t.Logf("start app %v", s3)

	trie.Insert("app")
	s3 = trie.Search("app") // 返回 True
	t.Logf("start app %v", s3)
}

func TestTrie_Search(t *testing.T) {
	type fields struct {
		children map[rune]*Trie
		isEnd    bool
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Trie{
				children: tt.fields.children,
				isEnd:    tt.fields.isEnd,
			}
			if got := this.Search(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
