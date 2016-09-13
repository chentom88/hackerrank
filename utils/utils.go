package utils

import (
	"fmt"
)

func ScanIntoArray(array []int) {
	arraySize := len(array)
	valInterfaces := make([]interface{}, arraySize)

	for i := 0; i < arraySize; i++ {
		valInterfaces[i] = &array[i]
	}

	fmt.Scanln(valInterfaces...)
}
