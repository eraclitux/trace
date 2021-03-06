// Copyright (c) 2015 Andrea Masi. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE.txt file.

// +build !debug

// Package trace exposes simple tracing capabilities.
//
// The idea is that during early stages of code development one wants
// a simple (and dirty :)) way to inspect values of vars
// that can be quickly disabled.
//
// To enable trace functions to print, just build, run or test with ``debug``
// tag otherwise they will be noops:
// 	go build -tags debug
// 	go test -tags debug
// 	go run -tags debug main.go
// stderr is used to not perturb Examples() functions.
//
// Once trace is no more needed, the provided command
// can be used to remove its presence in a code base,
// for example:
//	trace-clean -r .
// will recursively remove import and function calls from
// all go files in the current directory and all its subdirs.
//
// Credits
//
// Original idea by Dave Cheney http://dave.cheney.net.
package trace

// Println prints to stderr if '-tags debug'
// is used when building/running, noop otherwise.
// stderr is used to not perturb example tests.
func Println(args ...interface{}) {}

// Printf prints to stderr using supplied format
// if '-tags debug' is used when building/running,
// otherwise it is a no-op.
// stderr is used to not perturb example tests.
func Printf(format string, a ...interface{}) {}

// PrettyPrint prints types in a human readable
// format to stderr.
func PrettyPrint(name string, s interface{}) {}
