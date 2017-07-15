//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "strconv"
    "bytes"
    "os"
)

// Spiral is a square, 2D slice containing the numbers within the range of a particular number
// squared organized in a spiral fashion.
type Spiral [][]int

// NewSpiral returns a pointer to a Spiral made with the specified number.
func NewSpiral(n int) Spiral {

    // Make a new square, 2D grid.
    s := make(Spiral, n)
    for i := range s {
        s[i] = make([]int, n)
    }

    // Define the initial amount of steps to take when walking the grid in a spiral fashion.
    steps := map[string]int {
        "Right" : n - 1,
        "Down" : n - 2,
        "Left": n - 2,
        "Up" : n - 3,
    }

    // Starting point will be the top-left corner (0,0).
    pos := struct {
        x, y int
    }{
        0,
        0,
    }

    // Always start to fill the grid with 1.
    fill := 1

    // Start by moving toward the right then down, then left, then up, etc.
    dir := "Right"

    // Keep track of the number of steps we've taken in each dir.
    dist := 0

    // We've reached the end of the line, turn.
    turn := func(d string) {
        steps[dir] -= 2
        dist = 0
        dir = d
    }

    // Keep walking down the line.
    walk := func() bool {
        if dist == steps[dir] {
            return false
        }
        dist++
        return true
    }

    for fill <= (n * n) { // While we haven't reached the final number...

        s[pos.x][pos.y] = fill

        switch dir {
        case "Right":
            if walk() {
                pos.x++
                break
            }
            pos.y++
            turn("Down")
        case "Down":
            if walk() {
                pos.y++
                break
            }
            pos.x--
            turn("Left")
        case "Left":
            if walk() {
                pos.x--
                break
            }
            pos.y--
            turn("Up")
        case "Up":
            if walk() {
                pos.y--
                break
            }
            pos.x++
            turn("Right")
        }

        fill++

    }

    return s

}

func (s Spiral) String() string {

    var ns bytes.Buffer

    w := len(strconv.Itoa(len(s) * len(s)))

    for i := range s {
        for j := range s[i] {
            e := bytes.NewBufferString(strconv.Itoa(s[j][i]))
            for p := len(e.String()); p <= w; p++ {
                e.WriteString(" ")
            }
            ns.WriteString(e.String())
        }
        if i != len(s) - 1 {
            ns.WriteString("\n")
        }
    }

    return ns.String()

}

func main() {

    num, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Incorrect usage")
        os.Exit(1)
    }

    s := NewSpiral(num)
    fmt.Println(s)

}
