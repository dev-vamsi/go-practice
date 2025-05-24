package main

import (
	"flag"
	"fmt"
	"os"
)

func Args() {
	fmt.Println(os.Args)
}

func Flags() {
	addFlag := flag.String("add", "", "has that string to append")
	flag.Parse()
	fmt.Println(*addFlag)
}
