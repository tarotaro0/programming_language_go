package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type dupInfo struct {
	count int
	// A sturct{} takes up no space.
	// This is useful for HashSet.
	filenames map[string]struct{}
}

func (di dupInfo) getFileNames() string {
	var fns []string
	for fn := range di.filenames {
		fns = append(fns, fn)
	}
	return strings.Join(fns, ", ")
}

func main() {
	counts := make(map[string]*dupInfo)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		fmt.Fprint(os.Stderr, "no filename")
		return
	}

	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			if counts[line] == nil {
				// map is initialied as nil so must initialize yourself
				counts[line] = &dupInfo{filenames: map[string]struct{}{}}
			}
			counts[line].count++
			counts[line].filenames[filename] = struct{}{}
		}
	}

	for line, di := range counts {
		if di.count > 1 {
			fmt.Println("----------------")
			fmt.Printf("Duplicate String: %s\nDuplicate Count: %d\nDuplicate File: %v\n", line, di.count, di.getFileNames())
		}
	}
}
