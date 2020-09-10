package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/calebcase/base58"
)

func cannot(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		panic(err)
	}
}

func main() {
	var decode bool

	flag.BoolVar(&decode, "d", false, "decode")
	flag.Parse()

	input, err := ioutil.ReadAll(os.Stdin)
	cannot(err)

	if decode {
		dst, err := base58.Decode(string(input))
		cannot(err)

		os.Stdout.Write(dst)
	} else {
		dst := base58.Encode(input)

		os.Stdout.WriteString(dst)
	}
}
