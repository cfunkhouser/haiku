// Package haiku provides word lists and functions to generate haiku-style
// names. This package uses math/random, so be sure to seed well before using.
// Or don't. Just be aware.
package haiku

import (
	"math/rand"
	"strconv"
	"strings"
)

// Words to be used for naming something.
type Words []string

// Random word from amongst the Words.
func (l Words) Random() string {
	return l[rand.Intn(len(l))]
}

type haiku struct {
	seperator  string
	components []string
}

func (h haiku) String() string {
	return strings.Join(h.components, h.seperator)
}

// Option of a haiku-style name, to be passed to Custom.
type Option func(*haiku)

// WithSeperator to divide the segments of the haiku. If this is provided more
// than once, the last one provided wins.
func WithSeperator(sep string) Option {
	return func(h *haiku) {
		h.seperator = sep
	}
}

// Component of a haiku name. Should be passed to WithComponents.
type Component func() string

// Adjective component of a haiku name.
func Adjective() string {
	return Adjectives.Random()
}

// Noun component of a haiku name.
func Noun() string {
	return Nouns.Random()
}

const maxInt = 255

// Number component of a haiku name.
func Number() string {
	return strconv.Itoa(rand.Intn(maxInt))
}

// WithComponents to divide the segments of the haiku. If this is provided more
// than once, the last one provided wins.
func WithComponents(cs ...Component) Option {
	return func(h *haiku) {
		for _, c := range cs {
			h.components = append(h.components, c())
		}
	}
}

// DefaultSeperator for haiku names.
const DefaultSeperator = "-"

var defaultHaiku = haiku{
	seperator: DefaultSeperator,
}

// Custom haiku-style name.
func Custom(opts ...Option) string {
	h := defaultHaiku
	for _, o := range opts {
		o(&h)
	}
	return h.String()
}

// Simple haiku-style name.
func Simple() string {
	return Custom(WithComponents(Adjective, Noun))
}
