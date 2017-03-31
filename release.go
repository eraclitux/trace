// Copyright (c) 2015 Andrea Masi. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE.txt file.

// +build !debug

// Package trace is a simple tracing package.
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
// Credits
//
// Original idea by Dave Cheney http://dave.cheney.net.
package trace

// Traceln prints to stderr if '-tags debug'
// is used when building/running, noop otherwise.
// stderr is used to not perturb example tests.
func Traceln(args ...interface{}) {}

// Tracef prints to stderr using supplied format
// if '-tags debug' is used when building/running,
// noop otherwise.
// stderr is used to not perturb example tests.
func Tracef(format string, a ...interface{}) {}

// PrettyStruct prints struct in a human readable
// format to stderr.
func PrettyStruct(name string, s interface{}) {}
