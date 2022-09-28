package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args {
		if strings.EqualFold(arg, "Assignment3.txt") {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Println(err)
			}

			io.Copy(os.Stdout, f)

			break
		}
	}
}

func init() {
	ioutil.WriteFile("Assignment3.txt", []byte("Hello There!"), 0775)
}
