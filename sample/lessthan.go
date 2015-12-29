package main

import (
	"fmt"

	"github.com/ushios/versiongo"
)

func main() {
	current, _ := versiongo.NewVersion("1.0.0", versiongo.FuzzySplit)
	next, _ := versiongo.NewVersion("1.0.1", versiongo.FuzzySplit)

	result, err := versiongo.Compare(current, next)

	if err != nil {
		panic(err)
	}

	if versiongo.LessThan == result {
		fmt.Println("Need to update!!")
	}
}
