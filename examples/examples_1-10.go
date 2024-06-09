package examples

import (
	"fmt"
	"os"
	"strings"
)

// Prints command line arguments
func Echo1() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

func Echo2() {
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func Echo3() {
	var s, sep string

	sep = " "

	s = strings.Join(os.Args[1:], sep)

	fmt.Println(s)
}
