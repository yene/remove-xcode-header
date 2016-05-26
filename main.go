// Remove the generated Xcode comment header from .h, .m, .swift files recursive.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	if strings.HasPrefix(path, ".") || strings.Contains(path, "Carthage") {
		return nil
	}

	if strings.HasSuffix(path, ".h") || strings.HasSuffix(path, ".m") || strings.HasSuffix(path, ".swift") {
		removeXcodeHeader(path)
	}

	return nil
}

func removeXcodeHeader(path string) {
	dat, err := ioutil.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	if lines[0] == "//" && lines[3] == "//" && lines[6] == "//" {
		fmt.Printf("%s\n", path)
		//fmt.Println("Found an empty comment on first line, lets scrub.")

		linesToRemove := 7
		if len(lines[7]) == 0 {
			//fmt.Println("also remove the next empty line")
			linesToRemove = 8
		} else {
			fmt.Println(lines[8])
		}

		content := strings.Join(lines[linesToRemove:], "\n")
		err := ioutil.WriteFile(path, []byte(content), 0644)
		check(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a folder with .h, .m, .swift files.")
		os.Exit(1)
	}
	var path string = os.Args[1]
	err := filepath.Walk(path, visit)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
