package pkg3

import "fmt"

// Name package to print example
const (
	Name = "pkg 3"
)

// Print test
func Print(from, me string) {
	fmt.Println("Hello word iam function ", me)
	fmt.Println("Iam called from ", from)
}
