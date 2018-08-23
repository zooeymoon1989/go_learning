package main

import "fmt"

func recursion(n uint) uint {

	if n < 2 {
		return n
	}

	return recursion(n-2) + recursion(n-1)

}

func main() {
	n := uint(10)
	i := uint(0)

	for i < n {
		fmt.Print(recursion(i))
		i++
	}
}
