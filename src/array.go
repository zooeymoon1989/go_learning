package main

import "fmt"

var array = [...]uint8{1,2,3,4,5,6,7,8,9}

func main() {

	i := 0

	for i < len(array){
		fmt.Println(array[i])
		i ++
	}


}
