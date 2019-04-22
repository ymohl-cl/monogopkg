// Example1 est le premier exemple pour expliquer le context
package main

import (
	"github.com/ymohl-cl/monogopkg/pkg/pkg1"
	"github.com/ymohl-cl/monogopkg/pkg/pkg2"
	"github.com/ymohl-cl/monogopkg/pkg/pkg3"
	"github.com/ymohl-cl/monogopkg/pkg/pkg4"
	"github.com/ymohl-cl/monogopkg/pkg/pkg5"
)

func main() {
	pkg1.Print(pkg1.Name)
	pkg2.Print(pkg2.Name)
	pkg3.Print(pkg3.Name)
	pkg4.Print(pkg4.Name)
	pkg5.Print(pkg5.Name)
}
