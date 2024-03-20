package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateString(t *testing.T) {
	type TestCase struct {
		length int
	}
	testCases := []TestCase{
		{
			length: 5,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.length, len(GenerateString(tc.length)))
	}
}
