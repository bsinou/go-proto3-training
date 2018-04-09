package runtime

import (
	"fmt"
)

func traceFromOtherFile() {
	fmt.Println("In OtherFile.go, calling Trace from within the same package")
	Trace()
}
