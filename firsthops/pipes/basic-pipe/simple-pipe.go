// Canonical pattern to illustrate the use of the iuo.Pipe
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	pr, pw := io.Pipe()

	go func() {
		// in fact send an EOF signal on the pipe so that the reader knows we've reached end of file.
		defer pw.Close()
		if _, err := fmt.Fprintln(pw, "Hello"); err != nil {
			panic(fmt.Sprintf("cannot write: %v", err))
		}
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		panic(fmt.Sprintf("cannot read: %v", err))
	}
}
