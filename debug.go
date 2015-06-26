// Copyright (c) 2015 Andrea Masi. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE.txt file.

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

func init() {
	Traceln("Tracing enabled...")
}
