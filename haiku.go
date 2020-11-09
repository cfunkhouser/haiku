// Package haiku provides word lists and functions to generate haiku-style
// names. This package uses math/random, so be sure to seed well before using.
// Or don't. Just be aware.
package haiku

import (
	"fmt"
	"math/rand"
)

// Words to be used for naming something.
type Words []string

// Random word from amongst the Words.
func (l Words) Random() string {
	return l[rand.Intn(len(l))]
}

// Simple a haiku-style name.
func Simple() string {
	return fmt.Sprintf("%v-%v", Adjectives.Random(), Nouns.Random())
}

// SimpleWithNumber a haiku-style name.
func SimpleWithNumber() string {
	return fmt.Sprintf("%v-%v-%v", Adjectives.Random(), Nouns.Random(), rand.Intn(255))
}
