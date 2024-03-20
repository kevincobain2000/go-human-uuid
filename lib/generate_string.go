package lib

import (
	"math/rand"
	"time"

	"github.com/kevincobain2000/go-human-uuid/pkg"
)

// randomWhole generates a random whole number within the given range
func randomWhole(min, max int) int {
	// #nosec G404
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randGen.Intn(max-min) + min
}

// GenerateString generates a string based on the given length
func GenerateString(length int) string {
	var s string
	for i := 0; i < length; i++ {
		// EVEN places are in fact the odd placeholders in the string (as counting starts from 0), so they need to be consonants
		if i%2 == 0 {
			s += pkg.Consonants[randomWhole(0, len(pkg.Consonants))]
		} else {
			s += pkg.Vowels[randomWhole(0, len(pkg.Vowels))]
		}
	}
	return s
}
