package main

import (
	"fmt"

	"github.com/eraclitux/trace"
)

type dummy struct {
	Name string
}

func main() {
	trace.PrettyPrint("dummy:", dummy{"John"})
	trace.Printf("%s", "debug this!")
	fmt.Println("printed")
}
