package main

import (
	"DSA/pkg/arrays"
	"fmt"
)

func main() {
	fmt.Println("Main Started")
	if err := arrays.PrintArray([]int{1, 2, 3}); err != nil {
		fmt.Println(err.Error())
	}

}
