package main

import (
	"fmt"

	"github.com/eraclitux/trace"
)

type dummy struct {
	Name string
}

func main() {
	fmt.Println("printed")
	trace.PrettyPrint("dummy:", dummy{"John"})
}
