// Example2 defini les modules
package main

import (
	"github.com/ymohl-cl/monogopkg/pkg2/pkg1"
	"github.com/ymohl-cl/monogopkg/pkg2/pkg2"
	"github.com/ymohl-cl/monogopkg/pkg2/pkg3"
	"github.com/ymohl-cl/monogopkg/pkg2/pkg4"
	"github.com/ymohl-cl/monogopkg/pkg2/pkg5"
)

func main() {
	pkg1.Print(pkg1.Name)
	pkg2.Print(pkg2.Name)
	pkg3.Print(pkg3.Name)
	pkg4.Print(pkg4.Name)
	pkg5.Print(pkg5.Name)
}
