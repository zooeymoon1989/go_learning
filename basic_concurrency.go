package main

import (
	"io"
	"os"
	"time"
	"fmt"
)

func main() {
	go echo(os.Stdin , os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("time out")
	os.Exit(0)
}

func echo(in io.Reader , out io.Writer)  {
	io.Copy(out,in)
}