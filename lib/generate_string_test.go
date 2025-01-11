package lib

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randomWhole(minimum, maximum int) int {
	// #nosec G404
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randGen.Intn(maximum-minimum) + minimum
}

func TestGenerateString(t *testing.T) {
	for i := 0; i < 5; i++ {
		r := randomWhole(0, 100)
		s, _ := NewGenerator(
			func(opt *Options) error {
				opt.Length = r
				return nil
			},
		)
		n := s.Generate()
		assert.Equal(t, r, len(n))
	}
}
func TestRandomness(t *testing.T) {
	randoms := []string{}
	for i := 0; i < 1000; i++ {
		r := randomWhole(5, 100)
		s, _ := NewGenerator(
			func(opt *Options) error {
				opt.Length = r
				return nil
			},
		)
		// check if the string is already generated
		for _, v := range randoms {
			assert.NotEqual(t, v, s)
		}
		n := s.Generate()
		randoms = append(randoms, n)
		assert.Equal(t, r, len(n))
	}
}
