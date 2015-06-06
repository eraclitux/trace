// +build !debug

/*
	Package stracer is possibly the simplest tracing package.
	To enable its functions to print just build with ``debug`` tag otherwise they will do noop:
		go build -tags debug
	stderr is used to not perturb example functions.

	Credits

	Original idea is by Dave Cheney http://dave.cheney.net.

*/
package stracer

// Traceln prints to stderr if '-tags debug'
// is used when building, noop otherwise.
// stderr is used to not perturb example tests.
func Traceln(args ...interface{}) {}

// Tracef prints to stderr using supplied format
// if '-tags debug' is used when building,
// noop otherwise.
// stderr is used to not perturb example tests.
func Tracef(format string, a ...interface{}) {}
