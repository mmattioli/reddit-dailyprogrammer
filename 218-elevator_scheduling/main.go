//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
    "bufio"
)

type Direction string

const (
    Up Direction = "Up"
    Down Direction = "Down"
)

// A Request is a Rider's desire to go from one Floor (Origin) to another (Destination) at a
// particular point in time.
type Request struct {
    Time int
    Origin int
    Destination int
}

// A Rider uses an elevator Car to traverse between Floors. Riders can make several Requests to be
// transferred from one Floor to another.
type Rider struct {
    Id string
    Requests []Request
}

// When a Rider Arrives on a Floor, the first (active) Request is removed.
func (r *Rider) Arrive() {
    r.Requests = append(r.Requests[:0], r.Requests[1:]...)
}

// A Floor is a holding area for active Riders.
type Floor struct {
    Id int
    Ingress chan Rider
    Egress chan Rider
    Standby []Rider
    Time int
}

// Accept takes in Riders from an elevator Car. Upon Arrival, the Rider(s) will be moved to Standby
// if they have future Requests to be made.
func (f *Floor) Accept() {
    for len(f.Ingress) > 0 {
        r := <- f.Ingress
        if r.Arrive(); len(r.Requests) > 0 {
            f.Standby = append(f.Standby, r)
        }
    }
}

// Eject queues Riders to be picked up by an elevator Car.
func (f *Floor) Eject() {
    temp := []Rider{}
    for i := range f.Standby {
        switch {
        case f.Standby[i].Requests[0].Time <= f.Time:
            f.Egress <- f.Standby[i]
        default:
            temp = append(temp, f.Standby[i])
        }
    }
    f.Standby = temp
}

// An elevator Car moves between Floors to transport Riders.
type Car struct {
    Id string
    Capacity int
    Speed float64
    Position float64
    Passengers []Rider
}

// Pickup takes in as many Riders from a particular Floor as its Capacity will allow.
func (c *Car) Pickup(in <-chan Rider) {
    for len(in) > 0 && len(c.Passengers) < c.Capacity {
        r := <- in
        c.Passengers = append(c.Passengers, r)
    }
}

// Dropoff lets Riders out of the Car and onto the Floor.
func (c *Car) Dropoff(out chan<- Rider) {
    temp := []Rider{}
    for i := range c.Passengers {
        switch {
        case c.DoorsOpen(c.Passengers[i].Requests[0].Destination):
            out <- c.Passengers[i]
        default:
            temp = append(temp, c.Passengers[i])
        }
    }
    c.Passengers = temp
}

// DoorsOpen indicates whether or not the elevator Car is at the proper position to open its doors
// when it arrives at a particular Floor.
func (c *Car) DoorsOpen(f int) bool {
    return math.Abs(c.Position - float64(f)) < 0.001
}

// Move adjusts the elevator Car's Position depending on its Speed and the indicated Direction.
func (c *Car) Move(d Direction) {
    switch d {
    case Up:
        c.Position += c.Speed
    case Down:
        if c.Position > 1.0 {
            c.Position -= c.Speed
        }
    }
}

// RidersGoing indicates whether or not there are Riders in an elevator Car who desire to be
// transported to a Floor that is in the indicated Direction.
func (c *Car) RidersGoing(d Direction) bool {
    for i := range c.Passengers {
        switch d {
        case Up:
            if float64(c.Passengers[i].Requests[0].Destination) > c.Position {
                return true
            }
        case Down:
            if float64(c.Passengers[i].Requests[0].Destination) < c.Position {
                return true
            }
        }
    }
    return false
}

// ElevatorScheduling returns the number of "seconds" it would take to satisfy all Riders' Requests
// given pre-populated Floors and specific elevator Cars.
func ElevatorScheduling(c []Car, f []Floor) int {

    // Is there anyone still in the building?
    ridersAlive := func() bool {
        for i := range f {
            if (len(f[i].Egress) + len(f[i].Ingress) + len(f[i].Standby)) > 0 {
                return true
            }
        }
        for i := range c {
            if len(c[i].Passengers) > 0 {
                return true
            }
        }
        return false
    }

    // Is there anyone waiting to be picked up?
    ridersNotServiced := func(d Direction, fn int) bool {
        for i := range f {
            switch d {
            case Up:
                if f[i].Id > fn &&  len(f[i].Egress) > 0 {
                    return true
                }
            case Down:
                if f[i].Id < fn && len(f[i].Egress) > 0 {
                    return true
                }
            }
        }
        return false
    }

    var t int
    for t = 0; ridersAlive(); t++ {
        for i := range f {
                f[i].Time = t
                f[i].Accept()
                f[i].Eject()
        }
        for i := range c {
            for j := range f {
                if c[i].DoorsOpen(f[j].Id) {
                    c[i].Dropoff(f[j].Ingress)
                    c[i].Pickup(f[j].Egress)
                }
            }
            switch {
            case c[i].RidersGoing(Up):
                c[i].Move(Up)
            case c[i].RidersGoing(Down):
                c[i].Move(Down)
            case ridersNotServiced(Up, int(math.Trunc(c[i].Position))):
                c[i].Move(Up)
            case ridersNotServiced(Down, int(math.Trunc(c[i].Position))):
                c[i].Move(Down)
            case len(c[i].Passengers) == 0:
                c[i].Move(Down)
            }
        }
    }

    return t

}

func main() {

    parseInput := func(fn string) ([]Car, []Floor) {
        fh, err := os.Open(fn)
        if err != nil {
            log.Fatal(err)
        }
        defer fh.Close()

        var r []Rider
        var c []Car
        var f []Floor
        var fi int

        addRequestToExistingRider := func(id string, t, o, d int) bool {
            for i := range r {
                if id == r[i].Id {
                    r[i].Requests = append(r[i].Requests, Request{t, o, d})
                    return true
                }
            }
            return false
        }

        fs := bufio.NewScanner(fh)
        for fs.Scan() {
            flds := strings.Fields(fs.Text())
            switch {
            case strings.Contains(flds[0], "C"):
                cap, _ := strconv.Atoi(flds[1])
                spd, _ := strconv.ParseFloat(flds[2], 64)
                pos, _ := strconv.ParseFloat(flds[3], 64)
                c = append(c, Car{flds[0], cap, spd, pos, []Rider{}})
            case strings.Contains(flds[0], "R"):
                tme, _ := strconv.Atoi(flds[1])
                org, _ := strconv.Atoi(flds[2])
                dst, _ := strconv.Atoi(flds[3])
                if !addRequestToExistingRider(flds[0], tme, org, dst) {
                    r = append(r, Rider{flds[0], []Request{Request{tme, org, dst}}})
                }
                if dst > fi {
                    fi = dst
                }
            }
        }
        for i := 1; i <= fi; i++ {
            var s []Rider
            for j := range r {
                if r[j].Requests[0].Origin == i {
                    s = append(s, r[j])
                }
            }
            f = append(f, Floor{i, make(chan Rider, len(r)), make(chan Rider, len(r)), s, 0})
        }
        if err := fs.Err(); err != nil {
            log.Fatal(err)
        }
        return c, f
    }

    c, f := parseInput("challenge_input.txt")

    fmt.Printf("Took %ds to service all passengers\n", ElevatorScheduling(c, f))

}
