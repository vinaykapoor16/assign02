package main

import (
	"fmt"
	"strings"
)



func vinay(input string) string {
	return strings.ToUpper(input)
}

func factorial(n int) int {
    if n <= 0 {
        return 1
    }
    return n * factorial(n-1)
}

func isEven(number int) bool {
    return number%2 == 0
}

func fibonacci(n int) []int {
    sequence := make([]int, n)
    if n <= 0 {
        return sequence
    }

    sequence[0], sequence[1] = 0, 1
    for i := 2; i < n; i++ {
        sequence[i] = sequence[i-1] + sequence[i-2]
    }
    return sequence
}

func isPalindrome(str string) bool {
   
    cleanedStr := strings.ReplaceAll(strings.ToLower(str), " ", "")

    
    for i, j := 0, len(cleanedStr)-1; i < j; {
        if cleanedStr[i] != cleanedStr[j] {
            return false
        }
        i++
        j--
    }

    return true
}



func main() {
	lowercaseString := "vinay kapoor"
	uppercaseString := vinay(lowercaseString)
	fmt.Println(uppercaseString)

	result := factorial(5)
   	fmt.Printf("Factorial of 5 is %d\n", result)

	  num := 6
    if isEven(num) {
        fmt.Printf("%d is even.\n", num)
    } else {
        fmt.Printf("%d is odd.\n", num)
    }
 numTerms := 10
    fibSeq := fibonacci(numTerms)
    fmt.Printf("Fibonacci Sequence up to %d terms: %v\n", numTerms, fibSeq)

	inputStr := "A man a plan a canal Panama"
    if isPalindrome(inputStr) {
        fmt.Printf("\"%s\" is a palindrome.\n", inputStr)
    } else {
        fmt.Printf("\"%s\" is not a palindrome.\n", inputStr)
    }

}
