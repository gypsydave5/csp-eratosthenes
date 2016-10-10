package cspEratosthenes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEratosthenes(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	result := Eratosthenes(10)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Expected", expected, "but got", result)
	}
}

func ExampleEratosthenes() {
	result := Eratosthenes(10)
	fmt.Println(result)
	// Output: [2 3 5 7 11 13 17 19 23 29]
}
