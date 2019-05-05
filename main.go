// Idea from: https://github.com/etcd-io/etcd/blob/master/main.go
//
// Package main is a simple wrapper of the real balancedbracesd entrypoint package
// (located at fuchsberger.email/balancedbraces/balancedbracesdmain) to ensure that balancedbraces can become
// "go getable"; e.g. `go get fuchsberger.email/balancedbraces` could work as expected and
// builds a binary in $GOBIN/balancedbraces
//
// This package should NOT be extended or modified in any way; to modify the
// binary, work in the `balancedbraces/balancedbracesdmain` package.
//

package main

import "fuchsberger.email/balancedbracessrv/balancedbracesmain"

// depencency extracted for unit test
var mainFunc = balancedbracesmain.Main

func main() {
	mainFunc()
}
