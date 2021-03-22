// this is a tiny test tool to quickly validate the parser
package main

import (
	"bufio"
	"fmt"
	"os"

	dpkg "github.com/izkasi/golinpkg"
)

// const statusfile string = "/var/lib/dpkg/status"
const statusfile string = "../tests/status"

func main() {

	fmt.Println("Starting apt status parser for", statusfile)

	fileHandle, _ := os.Open(statusfile)
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	pkgs := dpkg.ParseStatus(fileScanner)

	fmt.Println(len(pkgs))

	fmt.Println(len(pkgs[9].Conffiles))

	fmt.Println(pkgs[9].Conffiles)

	fmt.Println(len(pkgs[9].Description))

	for x := range pkgs[9].Description {
		fmt.Println(x, ":", pkgs[9].Description[x])
	}

}
