package main

import (
	"fmt"
	"os"
)

func main() {
	if err := internal.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
