package main

import (
	"fmt"

	"github.com/eraclitux/trace"
)

func main() {
	fmt.Println("printed")

	trace.Printf("%s", "debug this!")

}
