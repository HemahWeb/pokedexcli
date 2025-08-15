package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "  hello  world  ", expected: []string{"hello", "world"}},
		{input: "\tHello\nWorld\t", expected: []string{"hello", "world"}},
		{input: "Hello,   world!", expected: []string{"hello,", "world!"}},
		{input: "\u00A0hello\u2003world\u00A0", expected: []string{"hello", "world"}}, // NBSP and em-space
		{input: "   \t\n ", expected: []string{}},
		{input: "", expected: []string{}},
		{input: "POKEDEX", expected: []string{"pokedex"}},
	}

	for _, c := range cases {
		cleaned := cleanInput(c.input)
		if !reflect.DeepEqual(cleaned, c.expected) {
			t.Errorf("cleanInput(%q) = %v; want %v", c.input, cleaned, c.expected)
		}
	}
}
