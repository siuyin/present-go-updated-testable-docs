// Package mypkg implements silly functions.
package mypkg

import "fmt"

// AFunc implements function a.
func AFunc() {
	fmt.Println("This is AFunc in mypkg")
}

// App implementes THE application.
func App() {
	aFunc()
	bFunc()
}
func aFunc() {}
func bFunc() {}
