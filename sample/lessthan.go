package main

import (
	"fmt"

	"github.com/ushios/versiongo"
)

func main() {
	current := "1.0.0"
	next := "1.0.1"

	result, err := versiongo.Compare(current, next)

	if err != nil {
		panic(err)
	}

	if versiongo.LessThan == result {
		fmt.Println("Need to update!!")
	}
}
