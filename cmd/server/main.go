package main

import (
	"fmt"
	"github.com/cuongnd9/go-grpc/pkg"
	"os"
)

func main() {
	if err := pkg.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
