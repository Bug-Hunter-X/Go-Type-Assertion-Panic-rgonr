# Go Type Assertion Panic

This repository demonstrates a common error in Go related to type assertions.  When type asserting an interface{}, if the underlying type does not match the target type, the program will panic. This example highlights this potential issue and provides a solution using error handling to prevent runtime panics.

## Bug

The `bug.go` file contains code that attempts to type assert an `interface{}` to `int32`.  If the interface{} does not hold an `int32`, this will cause a panic.

## Solution

The `bugSolution.go` file shows how to use a type switch to safely handle different underlying types of the interface, preventing panics.
