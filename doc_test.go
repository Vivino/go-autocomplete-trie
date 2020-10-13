package trie

import "fmt"

func Example() {
	t := New()
	t.Insert("Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	results := t.SearchAll("wdn")
	fmt.Println(results)

	results2 := t.SearchAll("tsd")
	fmt.Println(results2)

	// Output:
	// [Wednesday]
	// [Thursday, Tuesday, Wednesday]
	//
	// The second result demonstrates both fuzzy matching and the default Levenshtein distance scheme. This is configurable.
}

func Example_noFeatures() {
	t := New().CaseSensitive().WithoutFuzzy().WithoutLevenshtein().WithoutNormalisation()
	t.Insert("Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	results := t.SearchAll("t")
	fmt.Println(results)

	results2 := t.SearchAll("T")
	fmt.Println(results2)

	// Output:
	// []
	// [Tuesday, Thursday]
}
