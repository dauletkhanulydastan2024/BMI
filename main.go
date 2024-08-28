package main

import (
	"fmt"
)

type point struct {
	x, y string
}

func main() {
	p := point{"Das", "Beka"}
	fmt.Printf("struct1: %s", p)

}
