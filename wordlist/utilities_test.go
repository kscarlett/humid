package wordlist_test

import (
	"errors"
	"strings"
	"unicode"
)

// duplicate tests if the input contains a duplicate string
func duplicate(input []string) (err error, problems map[string]int) {
	words := make(map[string]int)

	for _, val := range input {
		words[val] += 1
	}

	problems = make(map[string]int)

	for word, count := range words {
		if count > 1 {
			problems[word] = count
			err = errors.New("duplicates found")
		}
	}

	return
}

// empty tests if the input is empty
func empty(input []string) error {
	if len(input) == 0 {
		return errors.New("given list is empty")
	}

	return nil
}

// ilegalCharacters checks whether there are any characters other than letters, numbers and spaces
func illegalCharacters(input []string) (err error, problems []string) {
	illegalChars := "-_`~!@#$%^&*()[]{}=+|\\'\"\n\t"

	for _, word := range input {
		if strings.ContainsAny(word, illegalChars) {
			problems = append(problems, word)
		}
	}

	return
}

// lowerCase tests if everything is lowercase
func lowercase(input []string) (err error, problems []string) {
	for _, word := range input {
		for _, char := range word {
			if unicode.IsLetter(rune(char)) && !unicode.IsLower(rune(char)) {
				problems = append(problems, word)
				err = errors.New("non-lowercase character found")
				break
			}
		}
	}

	return
}
