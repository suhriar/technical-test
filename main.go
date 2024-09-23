package main

import (
	"fmt"
	"strings"
)

// 1
func isPalindrome(s string) bool {
	original := strings.ToLower(s)
	reversed := []rune(original)
	length := len(reversed)

	for i := 0; i < length/2; i++ {
		reversed[i], reversed[length-1-i] = reversed[length-1-i], reversed[i]
	}

	return original == string(reversed)
}

// 2
func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxValue := arr[0]
	for _, num := range arr {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}

// 3
func printTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 4
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// 1
	fmt.Println("--- 1 ---")
	fmt.Println(isPalindrome("radar"))
	fmt.Println(isPalindrome("hello"))

	// 2
	fmt.Println("--- 2 ---")
	arr := []int{3, 5, 1, 9, 2}
	fmt.Println(findMax(arr))

	// 3
	fmt.Println("--- 3 ---")
	printTriangle(5)

	// 4
	fmt.Println("--- 4 ---")
	fmt.Println(factorial(5))
}
