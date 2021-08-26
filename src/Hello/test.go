package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result

	}
	return 1
}
func main() {
	fmt.Println("enter main")
	var n uint64 = 15
	var test_n = Factorial(n)
	fmt.Println(test_n)

}
