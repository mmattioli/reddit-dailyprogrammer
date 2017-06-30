//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "strconv"
    "os"
)

// NewSpiral returns a square, 2D slice containing the numbers within the range of a particular
// number, n, squared organized in a spiral fashion.
func NewSpiral(n int) [][]int {

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
    currentPosition := struct {
        x, y int
    }{
        0,
        0,
    }

    // Always start to fill the grid with 1.
    currentNumber := 1

    // Start by moving toward the right then down, then left, then up, etc.
    nextStep := "Right"

    // Keep track of the number of steps we've taken in each direction.
    stepsTaken := 0

    for currentNumber <= (n * n) { // While we haven't reached the final number...

        grid[currentPosition.x][currentPosition.y] = currentNumber

        turn := func(d string) {
            steps[nextStep] -= 2
            stepsTaken = 0
            nextStep = d
        }

        switch nextStep {
        case "Right":
            if stepsTaken < steps[nextStep] {
                currentPosition.x++
                stepsTaken++
                break
            }
            currentPosition.y++
            turn("Down")
        case "Down":
            if stepsTaken < steps[nextStep] {
                currentPosition.y++
                stepsTaken++
                break
            }
            currentPosition.x--
            turn("Left")
        case "Left":
            if stepsTaken < steps[nextStep] {
                currentPosition.x--
                stepsTaken++
                break
            }
            currentPosition.y--
            turn("Up")
        case "Up":
            if stepsTaken < steps[nextStep] {
                currentPosition.y--
                stepsTaken++
                break
            }
            currentPosition.x++
            turn("Right")
        }

        currentNumber++

    }

    return grid

}

// SpiralPrinter prints the elements of a square, 2D slice on the screen in which each element is
// equally spaced; it is assumed the specified slice's elements are arranged in a spiral fashion.
func SpiralPrinter(spiral [][]int) {

    elementWidth := len(strconv.Itoa(len(spiral) * len(spiral)))

    for i := range spiral {
        for j := range spiral[i] {
            element := strconv.Itoa(spiral[j][i])
            for padding := len(element); padding <= elementWidth; padding++ {
                element += " "
            }
            fmt.Printf("%s", element)
        }
        fmt.Printf("\n")
    }
}

func main() {

    specimen, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Incorrect usage")
        os.Exit(1)
    }

    spiral := NewSpiral(specimen)
    SpiralPrinter(spiral)

}
