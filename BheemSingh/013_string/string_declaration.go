package main

import "fmt"

// raw string literals (use of back quote ) and here is any charater may be print. multiple line string

var x string = `Hi My name is Bheem Singh Chouhan.
Here I am trying to learn golang.
best thing I heard about go is "GO is all about types"
`

// interpreted string literals (by use of double quotes) single line string

var y = " Hi this is golang playground. Play with go."

func main() {

	fmt.Println(x)

	fmt.Println(y)
}
