// Simply launch the mutexExample function using test framework
package examples_test

import (
    "testing"
    "github.com/bsinou/go-proto3-training/firsthops/examples"
)

func TestMutexExample(t *testing.T) {
	examples.MutexExample()
 }
