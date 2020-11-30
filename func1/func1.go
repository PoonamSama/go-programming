// A basic function
package main

import (
	"fmt"
)

func main() {
	foo("Poonam Gupta")
	fmt.Println("Hello, playground from main")
}

//func( r receivers) identifier (parameters){return(s)} {...}
func foo(s string) {
	fmt.Println("Hello,", s, "from foo")
}
