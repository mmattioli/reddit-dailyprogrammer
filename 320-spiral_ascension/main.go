//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "strconv"
    "os"
)

// Spiral is a square, 2D slice containing the numbers within the range of a particular number
// squared organized in a spiral fashion.
type Spiral struct {
    grid [][]int
}

// NewSpiral returns a pointer to a Spiral made with the specified number.
func NewSpiral(n int) *Spiral {

    // Make a new square, 2D grid.
    grid := make([][]int, n)
    for i := range grid {
        grid[i] = make([]int, n)
    }

    // Define the initial amount of steps to take when walking the grid in a spiral fashion.
    steps := map[string]int {
        "Right" : n - 1,
        "Down" : n - 2,
        "Left": n - 2,
        "Up" : n - 3,
    }

    // Starting point will be the top-left corner (0,0).
    position := struct {
        x, y int
    }{
        0,
        0,
    }

    // Always start to fill the grid with 1.
    fill := 1

    // Start by moving toward the right then down, then left, then up, etc.
    direction := "Right"

    // Keep track of the number of steps we've taken in each direction.
    distance := 0

    // We've reached the end of the line, turn.
    turn := func(d string) {
        steps[direction] -= 2
        distance = 0
        direction = d
    }

    // Keep walking down the line.
    walk := func() bool {
        if distance == steps[direction] {
            return false
        }
        distance++
        return true
    }

    for fill <= (n * n) { // While we haven't reached the final number...

        grid[position.x][position.y] = fill

        switch direction {
        case "Right":
            if walk() {
                position.x++
                break
            }
            position.y++
            turn("Down")
        case "Down":
            if walk() {
                position.y++
                break
            }
            position.x--
            turn("Left")
        case "Left":
            if walk() {
                position.x--
                break
            }
            position.y--
            turn("Up")
        case "Up":
            if walk() {
                position.y--
                break
            }
            position.x++
            turn("Right")
        }

        fill++

    }

    return &Spiral{grid}

}

func (s *Spiral) String() string {

    temp := ""

    width := len(strconv.Itoa(len(s.grid) * len(s.grid)))

    for i := range s.grid {
        for j := range s.grid[i] {
            element := strconv.Itoa(s.grid[j][i])
            for padding := len(element); padding <= width; padding++ {
                element += " "
            }
            temp += element
        }
        if i != len(s.grid) - 1 {
            temp += "\n"
        }
    }

    return temp

}

func main() {

    specimen, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Incorrect usage")
        os.Exit(1)
    }

    spiral := NewSpiral(specimen)
    fmt.Println(spiral)

}
