package arrays

import (
	"errors"
	"fmt"
)

// PrintArray prints the given array
func PrintArray(arr []int) error {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	return errors.New("some error")
}
