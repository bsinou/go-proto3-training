package examples

import (
	"fmt"

	"github.com/bsinou/go-proto3-training/firsthops/runtime"
)

// TraceFromOtherPackage is just a PoC while testing patterns with the runtime Package to make "bar" introspection
func TraceFromOtherPackage() {
	fmt.Println("In traceExample.go, calling Trace from another package")
	runtime.Trace()
}
