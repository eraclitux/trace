// Copyright (c) 2015 Andrea Masi. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE.txt file.

// +build debug

package trace

import (
	"encoding/json"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "[TRACE] ", log.Lmicroseconds)

func Traceln(args ...interface{}) {
	logger.Println(args...)
}

func Tracef(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func PrettyStruct(name string, s interface{}) {
	// TODO use reflection to get the name and remove arg?
	b, _ := json.MarshalIndent(s, "", "  ")
	logger.Println(name, "-", string(b))
}

func init() {
	Traceln("enabled!")
}
