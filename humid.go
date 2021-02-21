package humid

import (
	"strings"

	"github.com/kscarlett/humid/wordlist"
	rand "lukechampine.com/frand"
)

// Options defines a struct that passes in the generator options.
type Options struct {
	// List defines the wordlist to use
	List []string
	// AdjectiveCount defines the amount of adjectives to be used. Minimum value of 1.
	AdjectiveCount int
	// Separator defines the delimiter of the parts of the ID. This is usually a dash or underscore.
	Separator string
	// Capitalize sets whether to capitalize the first letter of each word.
	Capitalize bool
}

// Generate generates a new human id with the default options:
// - animals wordlist
// - 1 adjective
// - dash separator
// - no capitalization
func Generate() string {
	return GenerateWithOptions(&Options{
		List:           wordlist.Animals,
		AdjectiveCount: 1,
		Separator:      "-",
		Capitalize:     false,
	})
}

// GenerateWithOptions generates a new human id with the given set of options.
func GenerateWithOptions(opt *Options) string {
	len := opt.AdjectiveCount + 1
	str := make([]string, len)

	// Make sure adjectiveCount is at least 1 (lowest supported value)
	if opt.AdjectiveCount <= 0 {
		opt.AdjectiveCount = 1
	}

	// Get random strings
	for i := 0; i < opt.AdjectiveCount; i++ {
		str[i] = randomString(wordlist.Adjectives)
	}

	str[len-1] = randomString(opt.List)

	if opt.Capitalize {
		for n, v := range str {
			str[n] = strings.Title(v)
		}
	}

	id := strings.Join(str, opt.Separator)
	id = strings.ReplaceAll(id, " ", opt.Separator)

	return id
}

func randomString(arr []string) string {
	size := len(arr)

	if size == 0 {
		return ""
	}

	return arr[rand.Intn(size)]
}
