package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bfile := os.Args[1]

	fmt.Printf("BSON Convertor %s\n", bfile)
	dat, err := ioutil.ReadFile(bfile)
	if err != nil {
		fmt.Printf("Could not open file %s %s", bfile, err)
		os.Exit(-1)
	}

}
