package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollect(t *testing.T) {
	table := []struct {
		name     string
		trie     *Trie
		dict     []string
		search   string
		expected []string
	}{
		{
			name:   "Valid Prefix",
			dict:   []string{"hello"},
			trie:   New(),
			search: "h",
			expected: []string{
				"hello",
			},
		},
		{
			name:     "Invalid Prefix",
			dict:     []string{"hello"},
			trie:     New(),
			search:   "x",
			expected: []string{},
		},
		{
			name:   "With Fuzzy",
			dict:   []string{"hello"},
			trie:   New(),
			search: "elo",
			expected: []string{
				"hello",
			},
		},
		{
			name:     "Without Fuzzy",
			dict:     []string{"hello"},
			trie:     New().WithoutFuzzy(),
			search:   "elo",
			expected: []string{},
		},
		{
			name:   "With Normalisation To",
			dict:   []string{"héllö"},
			trie:   New(),
			search: "hello",
			expected: []string{
				"héllö",
			},
		},
		{
			name:   "With Normalisation From",
			dict:   []string{"hello"},
			trie:   New(),
			search: "héllö",
			expected: []string{
				"hello",
			},
		},
		{
			name:   "Case Insensitive To",
			dict:   []string{"HeLlO"},
			trie:   New(),
			search: "hello",
			expected: []string{
				"HeLlO",
			},
		},
		{
			name:   "Case Insensitive From",
			dict:   []string{"hello"},
			trie:   New(),
			search: "HeLlO",
			expected: []string{
				"hello",
			},
		},
		{
			name:     "Case Sensitive To",
			dict:     []string{"HeLlO"},
			trie:     New().CaseSensitive(),
			search:   "hello",
			expected: []string{},
		},
		{
			name:     "Case Sensitive From",
			dict:     []string{"hello"},
			trie:     New().CaseSensitive(),
			search:   "HeLlO",
			expected: []string{},
		},
		{
			name:   "Default levenshtein",
			dict:   []string{"hello"},
			trie:   New(),
			search: "hallo",
			expected: []string{
				"hello",
			},
		},
		{
			name:     "No levenshtein",
			dict:     []string{"hello"},
			trie:     New().WithoutLevenshtein(),
			search:   "hallo",
			expected: []string{},
		},
		{
			name: "Custom levenshtein",
			dict: []string{"hello"},
			trie: New().CustomLevenshtein(map[uint8]uint8{
				0:  0,
				10: 1,
				20: 2,
			}),
			search:   "hallo",
			expected: []string{},
		},
		//Problematic cases start here
		//This one will pass
		{
			name:     "Word order small => big",
			dict:     []string{"suspense", "suspenseful"},
			trie:     New(),
			search:   "susp",
			expected: []string{"suspense", "suspenseful"},
		},
		//This one won't
		{
			name:     "Word order big => small",
			dict:     []string{"suspenseful", "suspense"},
			trie:     New(),
			search:   "susp",
			expected: []string{"suspense", "suspenseful"},
		},
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			tt.trie.Insert(tt.dict...)
			actual := tt.trie.SearchAll(tt.search)
			assert.ElementsMatch(t, actual, tt.expected)
		})
	}
}

func BenchmarkInsert(b *testing.B) {
	t := New()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		t.Insert("hello")
	}
}

// prevent compiler optimization
var result []string

func BenchmarkSearch(b *testing.B) {
	t := New()
	t.Insert("hallo you")
	b.ReportAllocs()
	b.ResetTimer()
	var r []string
	for n := 0; n < b.N; n++ {
		r = t.Search("hello", 1)
	}
	result = r
}

func BenchmarkSearchAll(b *testing.B) {
	t := New()
	t.Insert("hallo you")
	b.ReportAllocs()
	b.ResetTimer()
	var r []string
	for n := 0; n < b.N; n++ {
		r = t.SearchAll("hello")
	}
	result = r
}
