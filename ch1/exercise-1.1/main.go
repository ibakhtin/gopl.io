// Copyright © 2021 Igor Bakhtin
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Exercise 1.1
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

//!-
