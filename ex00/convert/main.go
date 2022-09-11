package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/naka-c1024/Gopher_Dojo_module_00/ex00/convert/mypkg"
)

func main() {
	flag.Parse()
	if dirname := flag.Arg(0); dirname == "" {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		os.Exit(0)
	} else if flag.Arg(1) != "" {
		fmt.Fprintf(os.Stderr, "error: multiple arguments\n")
		os.Exit(0)
	} else {
		mypkg.MyWalk(dirname)
	}
}
