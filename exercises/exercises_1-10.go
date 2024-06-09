package exercises

import (
	"fmt"
	"os"
	"strings"
)

/* Modify the "Echo() example" to also print os.Args[0], the name of the command that invoked it */
func Echo1() {
	s := strings.Join(os.Args[:], " ")

	fmt.Println(s)
}

/* Modify the "Echo() example" to print the index and value of each of its arguments one per line */
func Echo2() {
	for i, arg := range os.Args[:] {
		fmt.Printf("%d --- %s\n", i, arg)
	}
}
