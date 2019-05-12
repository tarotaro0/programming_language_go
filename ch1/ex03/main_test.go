package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func getArgs() []string {
	var args []string
	n := 1000

	for i := 0; i < n; i++ {
		args = append(args, "a")
	}

	return args
}

func BenchmarkEx1(b *testing.B) {
	buf := new(bytes.Buffer)
	b.ResetTimer()

	s, sep := "", ""
	for _, arg := range getArgs() {
		s += sep + arg
		sep = " "
	}

	fmt.Fprintln(buf, s)
}

func BenchmarkStringsJoin(b *testing.B) {
	buf := new(bytes.Buffer)
	b.ResetTimer()
	fmt.Fprintln(buf, strings.Join(getArgs(), " "))
}
