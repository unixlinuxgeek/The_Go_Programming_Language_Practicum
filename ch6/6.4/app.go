package main

import (
	"fmt"
	"github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch6/6.4"
	"os"
)

func main() {
	c := []custom.Path{[]custom.Point{{1, 2}}, []custom.Point{{5, 11}}}
	var d []custom.Path = c[0:2]
	for _, s := range d {
		fmt.Fprintf(os.Stdout, "%g\n", s)
	}
}
