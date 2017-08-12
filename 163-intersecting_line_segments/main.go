//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "math"
)

// Point is a X,Y coordinate pair.
type Point struct {
    X float64
    Y float64
}

// LineSegment is a part of a line bound by two distinct Points.
type LineSegment struct {
    Id rune
    Start Point
    End Point
}

// Intersects reports whether there is an intersection with another arbitrary LineSegment.
func (ls1 LineSegment) Intersects(ls2 LineSegment) bool {

    // X coordinate of leftmost point of potential intersection.
    left := math.Max(math.Min(ls1.Start.X, ls1.End.X), math.Min(ls2.Start.X, ls2.End.X))

    // X coordinate of rightmost point of potential intersection.
    right := math.Min(math.Max(ls1.Start.X, ls1.End.X), math.Max(ls2.Start.X, ls2.End.X))

    // Y coordinate of highest point of potential intersection.
    top := math.Max(math.Min(ls1.Start.Y, ls1.End.Y), math.Min(ls2.Start.Y, ls2.End.Y))

    // Y coordinate of lowest point of potential intersection.
    bottom := math.Min(math.Max(ls1.Start.Y, ls1.End.Y), math.Max(ls2.Start.Y, ls2.End.Y))

    return top <= bottom && left <= right
    
}

func main() {
    ls := []LineSegment{
        LineSegment{'A', Point{-2.5, 0.5}, Point{3.5, 0.5}},
        LineSegment{'B', Point{-2.23, 99.99}, Point{-2.10, -56.23}},
        LineSegment{'C', Point{-1.23, 99.99}, Point{-1.10, -56.23}},
        LineSegment{'D', Point{100.1, 1000.34}, Point{2000.23, 2100.23}},
        LineSegment{'E', Point{1.5, -1.0}, Point{1.5, 1.0}},
        LineSegment{'F', Point{2.0, 2.0}, Point{3.0, 2.0}},
        LineSegment{'G', Point{2.5, 0.5}, Point{2.5, 2.0}},
    }
    for _, ls1 := range ls {
        var inters []rune
        for _, ls2 := range ls {
            if ls1.Intersects(ls2) && ls1 != ls2 {
                inters = append(inters, ls2.Id)
            }
        }
        fmt.Printf("%s: %s\n", string(ls1.Id), string(inters))
    }
}
