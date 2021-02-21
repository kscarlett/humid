package wordlist_test

import (
	"testing"

	"github.com/kscarlett/humid/wordlist"
)

func TestAdjectivesDuplicate(t *testing.T) {
	err, problems := duplicate(wordlist.Adjectives)

	if err != nil {
		t.Errorf("wordlist contains duplicates: %v", problems)
	}
}

func TestAdjectivesEmpty(t *testing.T) {
	err := empty(wordlist.Adjectives)

	if err != nil {
		t.Error("wordlist is empty")
	}
}

func TestAdjectivesIllegalCharacters(t *testing.T) {
	err, problems := illegalCharacters(wordlist.Adjectives)

	if err != nil {
		t.Errorf("wordlist contains illegal characters: %v", problems)
	}
}

func TestAdjectivesLowercase(t *testing.T) {
	err, problems := lowercase(wordlist.Adjectives)

	if err != nil {
		t.Errorf("wordlist contains non-lowercase letters: %v", problems)
	}
}
