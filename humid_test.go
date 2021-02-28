package humid_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/kscarlett/humid"
	"github.com/kscarlett/humid/wordlist"
)

// Examples
func ExampleGenerate() {
	fmt.Println(humid.Generate())
	// Example output: blue-hummingbird
}

func ExampleGenerateWithOptions() {
	fmt.Println(humid.GenerateWithOptions(&humid.Options{
		List:           wordlist.Adjectives,
		AdjectiveCount: 2,
		Separator:      "_",
		Capitalize:     true,
	}))
	// Example output: Bumpy_Brown_Cat
}

// Tests
func TestGenerateDefault(t *testing.T) {
	const expectedSeparator = "-"
	const expectedWordCount = 2

	result := humid.Generate()
	t.Logf("returned id: %s\n", result)

	// Test separator
	// Test word count
	words := strings.Split(result, expectedSeparator)

	if len(words) != expectedWordCount {
		t.Fatalf("expected %d words separated by \"%s\", got %d words instead", expectedWordCount, expectedSeparator, expectedWordCount)
	}

	// Test capitalization
	if !isLower(words[0]) {
		t.Errorf("expected word \"%s\" to be lowercase", words[0])
	}

	if !isLower(words[1]) {
		t.Errorf("expected word \"%s\" to be lowercase", words[1])
	}

	// Test wordlist
	adjectives := wordlist.Adjectives
	if !find(adjectives, words[0]) {
		t.Errorf("expected to find \"%s\" in wordlist it was not there", words[0])
	}

	if !find(adjectives, words[1]) {
		t.Errorf("expected to find \"%s\" in wordlist it was not there", words[1])
	}
}

func TestRandomQuality(t *testing.T) {
	id1 := humid.Generate()
	id2 := humid.Generate()

	if id1 == id2 {
		t.Errorf("random quality not as expected - %s and %s generated sequentially", id1, id2)
	}
}

// Benchmarks
func BenchmarkGenerateSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		humid.Generate()
	}
}

func BenchmarkGenerateWithOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		humid.GenerateWithOptions(&humid.Options{
			AdjectiveCount: 1,
			Separator:      "_",
			Capitalize:     false,
		})
	}
}

// Utility
func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func isTitle(s string) bool {
	if unicode.IsUpper(rune(s[0])) && unicode.IsLower(rune(s[1])) {
		return true
	}

	return false
}

func isLower(s string) bool {
	if unicode.IsLower(rune(s[0])) && unicode.IsLower(rune(s[1])) {
		return true
	}

	return false
}
