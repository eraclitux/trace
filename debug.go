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

func Println(a ...interface{}) {
	logger.Println(a...)
}

func Printf(format string, a ...interface{}) {
	logger.Printf(format, a...)
}

func PrettyPrint(name string, s interface{}) {
	// TODO use reflection to get the name and remove arg?
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		logger.Printf("error printing %q", name)
		return
	}
	logger.Println(name, "-", string(b))
}

func init() {
	Println("enabled!")
}
