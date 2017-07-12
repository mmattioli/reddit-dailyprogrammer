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

func ThreeSum(nums ...int) {

    // Calculate and log the total execution time.
    execTime := func(s time.Time) {
        log.Printf("Execution time: %s", time.Since(s))
    }

    defer execTime(time.Now())

    // Use a slice to store the triplets.
    var trpl [][3]int

    // Use a channel to send data across goroutines and a WaitGroup to ensure all goroutines finish
    // before closing the channel.
    c := make(chan [3]int)
    var wg sync.WaitGroup
    wg.Add(len(nums))

    // Keep track of all triplets found and ensure no duplicates.
    track := func(it <-chan [3]int) {
        go func() {
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
        }()
    }

    // Find all triplets whose sum equals zero (0).
    find := func(ot chan<- [3]int) {
        sort.Ints(nums)
        for i := range nums {
            go func(i int) {
                defer wg.Done()
                start, end := i + 1, len(nums) - 1
                for start < end {
                    val := nums[i] + nums[start] + nums[end]
                    switch {
                    case val == 0:
                        ot <- [3]int{nums[i], nums[start], nums[end]}
                        end--
                    case val > 0:
                        end--
                    default:
                        start++
                    }
                }
            }(i)
        }
        wg.Wait()
        close(c)
    }

    // Start the goroutines.
    track(c)
    find(c)

    // Display the triplets.
    for t := range trpl {
        fmt.Println(trpl[t])
    }

}

func main() {
    ThreeSum(4, 5, -1, -2, -7, 2, -5, -3, -7, -3, 1)
    ThreeSum(-1, -6, -3, -7, 5, -8, 2, -8, 1)
    ThreeSum(-5, -1, -4, 2, 9, -9, -6, -1, -7)
}
