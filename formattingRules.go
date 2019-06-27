package main

import (
	"fmt"
	"os"
)

/*
The exception to the “no-dead-code” rule is the blank identifier.  When you assign a value to the underscore character (“_“),
Go silently discards that value and does not complain about it being unused.  This is a common technique when you call a
function that returns multiple results, and you don’t actually need all of the results.

The os.State() function returns a FileInfo type for a given file path, and an error object that will be populated if no
file could be found.  In this case, we don’t actually care about the FileInfo return value.  We’re only interested in
whether or not the error return type gets populated.  So we assign the first of the two return types to the underscore
blank identifier, to prevent compile errors due to never using it.

READ more formatting rules here: https://steveperkins.com/go-for-java-programmers-formatting/
*/
func main() {
	path := "/tmp/test"
	if _, err := os.Stat("/tmp/test"); os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}
}
