package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	f, err := os.OpenFile("test.proto", os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	n, err := f.Write([]byte("Hello World!"))
	fmt.Println("n:", n)
	if err != nil {
		panic(err)
	}
}
