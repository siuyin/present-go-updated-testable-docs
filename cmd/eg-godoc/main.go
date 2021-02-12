package main

import (
	"fmt"

	"github.com/siuyin/present-go-updated-testable-docs/mypkg"
)

func main() {
	aFunc()
	mypkg.AFunc()
}

// aFunc implements function a.
func aFunc() {
	fmt.Println("This is function a")
}
