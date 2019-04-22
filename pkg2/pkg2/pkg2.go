package pkg2

import (
	"fmt"

	"github.com/ymohl-cl/monogopkg/pkg2/pkg3"
)

// Name pakage to print example
const (
	Name = "pkg 2"
)

// Print test
func Print(name string) {
	fmt.Println("Hello word iam function ", name)
	pkg3.Print(Name, pkg3.Name)
}
