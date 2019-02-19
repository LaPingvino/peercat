// Peercat is a tool for connecting random computers on the internet in a netcat like style
// using a shared connection string.
package main

import (
	"os"
	"io"
)

func main() {
	var (
		in io.Reader = os.Stdin
		out io.Writer = os.Stdout
	)
	if len(os.Args) > 1 {
		ins := make([]io.Reader, len(os.Args)-1)
		for i, file := range os.Args[1:] {
			f, err := os.Open(file)
			if err != nil {
				panic(err)
			}
			ins[i] = f
		}
		in = io.MultiReader(ins...)
	}
	io.Copy(out, in)
}
