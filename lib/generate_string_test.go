package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateString(t *testing.T) {
	for i := 0; i < 5; i++ {
		r := randomWhole(0, 100)
		s := GenerateString(r)
		assert.Equal(t, r, len(s))
	}
}
func TestRandomness(t *testing.T) {
	randoms := []string{}
	for i := 0; i < 1000; i++ {
		r := randomWhole(5, 100)
		s := GenerateString(r)
		// check if the string is already generated
		for _, v := range randoms {
			assert.NotEqual(t, v, s)
		}
		randoms = append(randoms, s)
		assert.Equal(t, r, len(s))
	}
}
