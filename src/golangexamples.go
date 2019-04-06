package main

import (
	"fmt"

	"github.com/ehteshamz/greetings"
)

func ConcatSlice(sliceToConcat []byte) string {
	result := ""
	for _, i := range sliceToConcat {
		result = result + string(i)
		result = result + string("-")
	}
	return result
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i, j := range sliceToEncrypt {
		sliceToEncrypt[i] = byte(int(j) + ceaserCount)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

func main() {
	inputt := []byte("AB")
	// result := ConcatSlice([]byte(inputt[:4]))
	// fmt.Printf(result+ "\n")

	Encrypt(inputt[:], 2)
	// fmt.Printf("%s", inputt)
	fmt.Println(EZGreetings("Usman"))
}
