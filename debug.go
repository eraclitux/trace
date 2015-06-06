// +build debug

package stracer

import (
	"fmt"
	"os"
)

func Traceln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func Tracef(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
