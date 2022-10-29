package main

import (
	"fmt"
	"github.com/tobgu/qframe"
	"strings"
)

func main() {
	input := `COL1,COL2
	a,1.5
	b,2.25
	c,3.0`

	f := qframe.ReadCSV(strings.NewReader(input))
	fmt.Println(f)
}
