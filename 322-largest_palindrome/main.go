//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "strconv"
    "bytes"
)

// LargestPalindrome returns the largest integer that is a palindrome and has two factors both of
// string length n.
func LargestPalindrome(n int) int {

    isPalindrome := func(n int) bool {
        ten := 1
        for t := n; t != 0; t /= 10 {
            ten *= 10
        }
        ten /= 10
        for n != 0 {
            if (n / ten) != (n % 10) {
                return false
            }
            n %= ten
            n /= 10
            ten /= 100
        }
        return true
    }

    startingFactor := func(n int) int {
        var b bytes.Buffer
        for i := 0; i < n; i++ {
            b.WriteString("9")
        }
        f, _ := strconv.Atoi(b.String())
        return f
    }

    sf := startingFactor(n)

    var f1, f2 int
    for i := sf; len(strconv.Itoa(i)) == n; i-- {
        for j := sf; len(strconv.Itoa(j)) == n; j-- {
            if isPalindrome(i * j) && (i * j) > (f1 * f2) {
                f1, f2 = i, j
            }
        }
    }

    return f1 * f2

}

func main() {
    fmt.Println(LargestPalindrome(1))
    fmt.Println(LargestPalindrome(2))
    fmt.Println(LargestPalindrome(3))
    fmt.Println(LargestPalindrome(4))
    fmt.Println(LargestPalindrome(5))
    fmt.Println(LargestPalindrome(6))
}
