package runtime

import (
	"fmt"
	"runtime"
)

// Trace implements a simple pattern to retrieve syslog information easily.
// thanks to https://stackoverflow.com/questions/25927660/golang-get-current-scope-of-function-name#25927915
func Trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	// fmt.Printf("I'm really here... %s:%d %s\n", file, line, f.Name())
	fmt.Printf("Calling method: %s, L.%d of %s\n", f.Name(), line, file)
}

// func main() {
// 	Trace()
// 	Trace()
// 	traceFromOtherFile()
// }
