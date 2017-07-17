//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
)

// It takes one (1) minute to travel from one stop to another, there are eight (8) hours in a work
// day and sixty (60) minutes in an hour which totals four-hundred and eighty (480) minutes and,
// therefore, four-hundred and eighty (480) trips made in a work day.
const MaxTrips int = 480

// A BusDriver has a daily route composed of MaxTrips stops and a collection of gossip which can be
// shared with other BusDrivers.
type BusDriver struct {
    DailyRoute []int
    Gossips []*BusDriver
}

// NewBusDriver creates and initializes a new BusDriver. The BusDriver's route for the entire day is
// filled out and starts with one (1) gossip to share.
func NewBusDriver(r ...int) *BusDriver {

    var bd BusDriver

    // Initial gossip.
    bd.Gossips = append(bd.Gossips, &bd)

    // Determine daily route.
    var t int
    for {
        for s := range r {
            if t == MaxTrips {
                return &bd
            }
            bd.DailyRoute = append(bd.DailyRoute, r[s])
            t++
        }
    }

}

// ExchangeGossip shares unknown gossip between two BusDrivers.
func ExchangeGossip(bd1, bd2 *BusDriver) {
    giveGossip := func(src, dst *BusDriver) {
        for sg := range src.Gossips {
            var known bool
            for dg := range dst.Gossips {
                if src.Gossips[sg] == dst.Gossips[dg] {
                    known = true
                }
            }
            if !known {
                dst.Gossips = append(dst.Gossips, src.Gossips[sg])
            }
        }
    }
    giveGossip(bd1, bd2)
    giveGossip(bd2, bd1)
}

// AllGossipExchanged determines if all of the BusDrivers have shared their gossip with each other.
func AllGossipExchanged(bds ...*BusDriver) bool {
    for i := 0; i < len(bds) - 1; i++ {
        switch {
        // If a BusDriver only has 1 gossip then hasn't receeived anything.
        case len(bds[i].Gossips) == 1:
            return false
        // If any two BusDriver's gossips are not of equal length then everyone has not shared their
        // gossip with everyone.
        case len(bds[i].Gossips) != len(bds[i + 1].Gossips):
            return false
        }
    }
    return true
}

// BusDriverGossipExchange calculates the number of stops each BusDriver must make before they all
// have shared each other's gossip. Returns -1 if all of the BusDrivers have not shared and heard
// all of the gossip there is to share and hear by the end of their routes.
func BusDriverGossipExchange(r ...[]int) int {

    var drvs []*BusDriver

    for br := range r {
        drvs = append(drvs, NewBusDriver(r[br]...))
    }

    for t := 0; t < MaxTrips; t++ {
        for src := range drvs {
            for dst := range drvs {
                switch {
                // Dont't exchange gossip with the same BusDriver.
                case drvs[src] == drvs[dst]:
                    continue
                // Two different BusDrivers at the same bus stop.
                case drvs[src].DailyRoute[t] == drvs[dst].DailyRoute[t]:
                    ExchangeGossip(drvs[src], drvs[dst])
                }
            }
        }
        if AllGossipExchanged(drvs...) {
            return t
        }
    }

    return -1

}

func main() {

    tests := [][][]int{
        [][]int {
            []int{3, 1, 2, 3},
            []int{3, 2, 3, 1},
            []int{4, 2, 3, 4, 5},
        },
        [][]int{
            []int{2, 1, 2},
            []int{5, 2, 8},
        },
        [][]int{
            []int{7, 11, 2, 2, 4, 8, 2, 2},
            []int{3, 0, 11, 8},
            []int{5, 11, 8, 10, 3, 11},
            []int{5, 9, 2, 5, 0, 3},
            []int{7, 4, 8, 2, 8, 1, 0, 5},
            []int{3, 6, 8, 9},
            []int{4, 2, 11, 3, 3},
        },
        [][]int {
            []int{12, 23, 15, 2, 8, 20, 21, 3, 23, 3, 27, 20, 0},
            []int{21, 14, 8, 20, 10, 0, 23, 3, 24, 23, 0, 19, 14, 12, 10, 9, 12, 12, 11, 6, 27, 5},
            []int{8, 18, 27, 10, 11, 22, 29, 23, 14},
            []int{13, 7, 14, 1, 9, 14, 16, 12, 0, 10, 13, 19, 16, 17},
            []int{24, 25, 21, 4, 6, 19, 1, 3, 26, 11, 22, 28, 14, 14, 27, 7, 20, 8, 7, 4, 1, 8, 10, 18, 21},
            []int{13, 20, 26, 22, 6, 5, 6, 23, 26, 2, 21, 16, 26, 24},
            []int{6, 7, 17, 2, 22, 23, 21},
            []int{23, 14, 22, 28, 10, 23, 7, 21, 3, 20, 24, 23, 8, 8, 21, 13, 15, 6, 9, 17, 27, 17, 13, 14},
            []int{23, 13, 1, 15, 5, 16, 7, 26, 22, 29, 17, 3, 14, 16, 16, 18, 6, 10, 3, 14, 10, 17, 27, 25},
            []int{25, 28, 5, 21, 8, 10, 27, 21, 23, 28, 7, 20, 6, 6, 9, 29, 27, 26, 24, 3, 12, 10, 21, 10, 12, 17},
            []int{26, 22, 26, 13, 10, 19, 3, 15, 2, 3, 25, 29, 25, 19, 19, 24, 1, 26, 22, 10, 17, 19, 28, 11, 22, 2, 13},
            []int{8, 4, 25, 15, 20, 9, 11, 3, 19},
            []int{24, 29, 4, 17, 2, 0, 8, 19, 11, 28, 13, 4, 16, 5, 15, 25, 16, 5, 6, 1, 0, 19, 7, 4, 6},
            []int{16, 25, 15, 17, 20, 27, 1, 11, 1, 18, 14, 23, 27, 25, 26, 17, 1},
        },
    }

    for t := range tests {
        fmt.Println(BusDriverGossipExchange(tests[t]...))
    }

}
