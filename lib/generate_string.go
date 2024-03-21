package lib

import (
	"math/rand"
	"time"

	"github.com/kevincobain2000/go-human-uuid/pkg"
)

const (
	defaultLength = 10
)

type Options struct {
	Length int
}

type Option func(*Options) error

func NewGenerator(opts ...Option) (*Options, error) {
	opt := &Options{}
	if opt.Length == 0 {
		opt.Length = defaultLength
	}

	for _, o := range opts {
		if err := o(opt); err != nil {
			return nil, err
		}
	}

	return opt, nil
}

// randomWhole generates a random whole number within the given range
func (o *Options) randomWhole(min, max int) int {
	// #nosec G404
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randGen.Intn(max-min) + min
}

// GenerateString generates a string based on the given length
func (o *Options) Generate() string {
	var s string
	for i := 0; i < o.Length; i++ {
		// EVEN places are in fact the odd placeholders in the string (as counting starts from 0), so they need to be consonants
		if i%2 == 0 {
			s += pkg.Consonants[o.randomWhole(0, len(pkg.Consonants))]
		} else {
			s += pkg.Vowels[o.randomWhole(0, len(pkg.Vowels))]
		}
	}
	return s
}
