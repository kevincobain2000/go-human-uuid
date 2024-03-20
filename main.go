package main

import (
	"fmt"

	"github.com/kevincobain2000/go-human-uuid/lib"
)

func GenerateString(length int) string {
	return lib.GenerateString(length)
}

func main() {
	for i := 0; i < 5; i++ {
		s := GenerateString(5)
		fmt.Println("GenerateString:", s)
	}
}
