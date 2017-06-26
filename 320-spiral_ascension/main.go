//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "strconv"
    "os"
)

// Position is a simple X,Y coordinate pair corresponding to a location on a 2D grid.
type Position struct {
    x, y int
}

// NewGrid returns a square, 2D slice; the size is determined by the span and each value is
// initialized to 0.
func NewGrid(span int) [][]int {
    grid := make([][]int, span)
    for i := range grid {
        grid[i] = make([]int, span)
    }
    return grid
}

// NewSpiral returns a square, 2D slice containing the numbers within the range of a particular
// number, n, squared organized in a spiral fashion.
func NewSpiral(n int) [][]int {

    // Make a new grid.
    grid := NewGrid(n)

    // Define the initial amount of steps to take when walking the grid in a spiral fashion.
    rightSteps := n - 1
    downSteps := n - 2
    leftSteps := n - 2
    upSteps := n - 3

    // Starting point will be the top-left corner (0,0).
    currentPosition := &Position{0, 0}

    // Always start to fill the grid with 1.
    currentNumber := 1

    // Start by moving toward the right then down, then left, then up, etc.
    nextStep := "Right"

    // Keep track o the number of steps we've taken in each direction.
    stepsTaken := 0

    for currentNumber <= (n * n) { // While we haven't reached the final number...
        grid[currentPosition.x][currentPosition.y] = currentNumber
        switch nextStep {
        case "Right":
            if stepsTaken < rightSteps {
                currentPosition.x++
                stepsTaken++
            } else {
                currentPosition.y++
                rightSteps -= 2
                stepsTaken = 0
                nextStep = "Down"
            }
        case "Down":
            if stepsTaken < downSteps {
                currentPosition.y++
                stepsTaken++
            } else {
                currentPosition.x--
                downSteps -= 2
                stepsTaken = 0
                nextStep = "Left"
            }
        case "Left":
            if stepsTaken < leftSteps {
                currentPosition.x--
                stepsTaken++
            } else {
                currentPosition.y--
                leftSteps -= 2
                stepsTaken = 0
                nextStep = "Up"
            }
        case "Up":
            if stepsTaken < upSteps {
                currentPosition.y--
                stepsTaken++
            } else {
                currentPosition.x++
                upSteps -= 2
                stepsTaken = 0
                nextStep = "Right"
            }
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
        line := ""
        for j := range spiral[i] {
            element := strconv.Itoa(spiral[j][i])
            for padding := len(element); padding <= elementWidth; padding++ {
                element += " "
            }
            line += element
        }
        fmt.Println(line)
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
