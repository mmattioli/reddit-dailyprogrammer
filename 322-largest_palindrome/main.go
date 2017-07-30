//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "time"
    "log"
    "math"
    "sync"
)

// LargestPalindrome returns the largest integer that is a palindrome and has two factors both of
// string length n.
func LargestPalindrome(n int) int {

    // Calculate and log the total execution time.
    defer func(t time.Time) {
        log.Printf("Execution time: %s", time.Since(t))
    }(time.Now())

    // Determine if an integer is a palindrome.
    isPalindrome := func(n int) bool {
        num, rev := n, 0
        for n > 0 {
            rev = (rev * 10) + (n % 10)
            n /= 10
        }
        return num == rev
    }

    // Maximum and minimum, respectively, given the input.
    maxFactor, minFactor := int(math.Pow10(n)) - 1, int(math.Pow10(n - 1))

    var f1, f2 int
    var wg sync.WaitGroup
    wg.Add(9 * minFactor)
    for i := maxFactor; i >= minFactor; i-- {
        go func(i int) {
            defer wg.Done()
            for j := maxFactor; j >= minFactor; j-- {
                if isPalindrome(i * j) && (i * j) > (f1 * f2) {
                    f1, f2 = i, j
                }
            }
        }(i)
    }
    wg.Wait()

    return f1 * f2

}

func main() {
    for n := 1; n <= 6; n++ {
        fmt.Printf("%d: %d\n", n, LargestPalindrome(n))
    }
}
