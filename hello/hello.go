// One of the first bits of Go code I ever wrote.
package main

import (
	"fmt"

	"go.jlucktay.dev/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("\n!oG ,olleH"))

	doForLoop()

	doMath()

	doSwitchOS()
	doSwitchTime()
	doSwitchTrue()

	deferHello()
	deferCount()
}
