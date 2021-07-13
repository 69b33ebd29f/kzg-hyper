package main

import (
	"fmt"

	gmcl "github.com/69b33ebd29f/mcl-wrapper"
	// "github.com/69b33ebd29f/kzg-hyper/ff"
)

func main() {
	fmt.Println("Hello, World!")
	gmcl.InitFromString("bls12-381")
}
