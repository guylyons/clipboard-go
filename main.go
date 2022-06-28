package main

import (
	"fmt"

	"golang.design/x/clipboard"
)

func main() {
	fmt.Println("Hey!")
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, []byte("text data"))

}
