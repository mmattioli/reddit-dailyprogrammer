//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "sort"
    "sync"
    "time"
    "log"
)

// Triplet is a slice containing three (3) integers whose sum equals zero (0).
type Triplet [3]int

// ThreeSum returns a slice of Triplets given a set of real numbers.
func ThreeSum(nums ...int) []Triplet {

    // Calculate and log the total execution time.
    defer func(t time.Time) {
        log.Printf("Execution time: %s", time.Since(t))
    }(time.Now())

    // Use a slice to store the triplets.
    var trpl []Triplet

    // Use a channel to send data across goroutines and a WaitGroup to ensure all goroutines finish
    // before closing the channel.
    c := make(chan Triplet)
    var wg sync.WaitGroup
    wg.Add(len(nums))

    // Keep track of all triplets found and ensure no duplicates.
    go func(it <-chan Triplet) {
        for i := range it {
            var found bool
            for t := range trpl {
                if trpl[t][0] == i[0] && trpl[t][1] == i[1] && trpl[t][2] == i[2] {
                    found = true
                }
            }
            if !found {
                trpl = append(trpl, i)
            }
        }
    }(c)

    // Find all triplets whose sum equals zero (0).
    sort.Ints(nums)
    for i := range nums {
        go func(i int, ot chan<- Triplet) {
            defer wg.Done()
            start, end := i + 1, len(nums) - 1
            for start < end {
                val := nums[i] + nums[start] + nums[end]
                switch {
                case val == 0:
                    ot <- Triplet{nums[i], nums[start], nums[end]}
                    end--
                case val > 0:
                    end--
                default:
                    start++
                }
            }
        }(i, c)
    }

    wg.Wait()
    return trpl

}

func main() {
    tests := [][]int{
        []int{4, 5, -1, -2, -7, 2, -5, -3, -7, -3, 1},
        []int{-1, -6, -3, -7, 5, -8, 2, -8, 1},
        []int{-5, -1, -4, 2, 9, -9, -6, -1, -7},
    }
    for t := range tests {
        fmt.Println(ThreeSum(tests[t]...))
    }
}
