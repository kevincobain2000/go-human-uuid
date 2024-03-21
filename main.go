package main

import (
	"fmt"

	"github.com/kevincobain2000/go-human-uuid/lib"
)

func GenerateString(length int) (string, error) {
	options := []lib.Option{
		func(opt *lib.Options) error {
			opt.Length = length
			return nil
		},
	}

	// Create generator with options
	// gen, err := lib.NewGenerator(options...)
	gen, err := lib.NewGenerator(options...)

	// Generate and print string
	return gen.Generate(), err
}

func main() {
	for i := 0; i < 5; i++ {
		s, _ := GenerateString(5)
		fmt.Println("GenerateString:", s)
	}
}
