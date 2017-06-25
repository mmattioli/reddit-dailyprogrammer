//
// Written by Michael Mattioli
//


package main

import (
    "fmt"
    "strconv"
    "os"
)

// Position is a simple X,Y coordinate pair corresponding to a location on a 2D grid
type Position struct {
    x, y int
}

// NewGrid returns a square 2D slice with each element initialized to 0
func NewGrid (specimen int) [][]int {
    grid := make([][]int, specimen)
    for i := range grid {
        grid[i] = make([]int, specimen)
    }
    return grid
}

// SpiralMaker returns a square grid containing the numbers within the range of a particular number
// squared organized in a spiral fashion
func SpiralMaker(s int) [][]int {

    // Make a new grid
    grid := NewGrid(s)

    // Define the initial amount of steps to take when walking the grid in a spiral fashion
    rightSteps := s - 1
    downSteps := s - 2
    leftSteps := s - 2
    upSteps := s - 3

    // Starting point will be the top-left corner (0,0)
    currentPosition := &Position{0, 0}

    // Always start to fill the grid with 1
    currentNumber := 1

    // Start by moving toward the right then down, then left, then up, etc.
    nextStep := "Right"

    // Keep track o the number of steps we've taken in each direction
    stepsTaken := 0

    for currentNumber <= (s * s) { // While we haven't reached the final number...
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

// SpiralPrinter prints a spiral grid on the screen in which each element is equally spaced
func SpiralPrinter (g [][]int) {

    elementWidth := len(strconv.Itoa(len(g) * len(g)))

    for i := range g {
        line := ""
        for j := range g[i] {
            element := strconv.Itoa(g[j][i])
            for p := len(element); p <= elementWidth; p++ {
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

    spiral := SpiralMaker(specimen)
    SpiralPrinter(spiral)

}
