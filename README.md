<h1 align="center">Go-Autocomplete-Trie</h1>

<div align="center">
  An <code>autocompl...</code> library for Go by Vivino.
</div>

<br />

## What Is it

Go-Autocomplete-Trie is a simple, configurable autocompletion library for Go. Simply build a dictionary with a slice of strings, optionally configure, and then search.

## How to Use

Make a default Trie like so: 

```trie := gat.New()``` 

The default Trie has *fuzzy* search enabled, string *normalisation* enabled, a default *levenshtein* scheme and is *case insensitive* by default.

Next, just add some strings to the dictionary.

```
trie.Insert("Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")
```

Next, search.

```
trie.SearchAll("wdn")

-> []string{"Wednesday"}
```

Levenshtein is enabled by default.

```
trie.SearchAll("urs")

-> []string{"Thursday", "Tuesday"}
```

To turn off the features...

```
trie.WithoutLevenshtein().WithoutNormalisation().WithoutFuzzy().CaseSensitive()
```

Now...

```
trie.SearchAll("urs")

-> []string{}

trie.SearchAll("Thu")

-> []string{"Thursday"}
```
