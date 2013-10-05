package main

import (
	"os"

	"fmt"
	"strings"

	"ac-get.darkdna.net/checker"
)

type res struct {
	indent int
}

func (_ res) BeginCheck() {
	// Do Nothing.
}

func (_ res) EndCheck() {
	// Do Nothing.
}

func (r res) Done() {
	// Do nothing.
}

func (r res) BeginTest(format string, args ...interface{}) checker.Results {
	fmt.Printf("%sBeginning: %s\n", strings.Repeat("  ", r.indent), fmt.Sprintf(format, args...))
	return res{r.indent + 1}
}

func (r res) Fail(format string, args ...interface{}) {
	fmt.Printf("%sFailed: %s\n", strings.Repeat("  ", r.indent), fmt.Sprintf(format, args...))
}

func (r res) Success() {
	fmt.Printf("%sSucceeded.\n", strings.Repeat("  ", r.indent))
}

func (r res) Warn(format string, args ...interface{}) {
	fmt.Printf("%sWarning: %s\n", strings.Repeat("  ", r.indent), fmt.Sprintf(format, args...))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: acg-checkup <repo-url>\n")

		return
	}

	r := res{0}

	checker.Run(os.Args[1], r)
}
