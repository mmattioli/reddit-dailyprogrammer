//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "reflect"
)

// A Tour is an activity that someone can partake in.
type Tour struct {
    Id string
    Name string
    Price float64
}

func (t Tour) String() string {
    return t.Id
}

// A SalesPromotion is a rule which determines the eligibility for price adjustments and/or free
// Tours.
type SalesPromotion func(t ...Tour) ([]Tour, float64)

// A ShoppingCart contains the Tours intended for purchase and any SalesPromotions that are to be
// applied to an order.
type ShoppingCart struct {
    Tours []Tour
    SalesPromotions []SalesPromotion
}

// AddSalesPromotion adds a SalesPromotion to a ShoppingCart.
func (s *ShoppingCart) AddSalesPromotion(sp ...SalesPromotion) {
    for i := range sp {
        var found bool
        for j := range s.SalesPromotions {
            if reflect.ValueOf(s.SalesPromotions[j]).Pointer() == reflect.ValueOf(sp[i]).Pointer() {
                found = true
            }
        }
        if !found {
            s.SalesPromotions = append(s.SalesPromotions, sp[i])
        }
    }
}

// AddTour adds a tour to a ShoppingCart.
func (s *ShoppingCart) AddTour(t ...Tour) {
    for i := range t {
        s.Tours = append(s.Tours, t[i])
    }
}

// Empty removes any/all Tours and SalesPromotions from a ShoppingCart.
func (s *ShoppingCart) Empty() {
    s.Tours = s.Tours[:0]
    s.SalesPromotions = s.SalesPromotions[:0]
}

// Review displays the Tours that were added for purchase and any/all adjustments that result from
// the SalesPromotions that are in effect.
func (s *ShoppingCart) Review() {
    var st, ap float64
    for i := range s.Tours {
        st += s.Tours[i].Price
    }
    var at []Tour
    for i := range s.SalesPromotions {
        t, a := s.SalesPromotions[i](s.Tours...)
        at = append(at, t...)
        ap += a
    }
    const summary = `
Items: %s
Subtotal: $%.2f
Sales promotion adjustments
    Tours added: %s Price adjustments: $%.2f
Grand total: $%.2f
`
    fmt.Printf(summary, s.Tours, st, at, ap, (st + ap))
}

func main() {

    OH := Tour{"OH", "Opera House Tour", 300.00}
    BC := Tour{"BC", "Sydney Bridge Climb", 110.00}
    SK := Tour{"SK", "Sydney Sky Tower", 30.00}

    // Purchasing three (3) Opera House tour yields a free Opera House tour.
    ThreeForTwoOperaHouse := func(t ...Tour) ([]Tour, float64) {
        var cnt int
        for i := range t {
            if t[i].Id == "OH" {
                cnt++
            }
        }
        return nil, float64(cnt / 3) * -OH.Price
    }

    // Purchasing one (1) Opera House tour yields a free Sky Tower tour.
    FreeSkyTowerWithOperaHouse := func(t ...Tour) ([]Tour, float64) {
        var cntOH, cntSK int
        for i := range t {
            switch t[i].Id {
            case "OH":
                cntOH++
            case "SK":
                cntSK++
            }
        }
        switch {
        // There are no Opera House tours being purchased so this promotion does not apply.
        case cntOH == 0:
            return nil, 0.00
        // Yield a credit for the same number of Sky Tower tours being purchased as the number of
        // Opera House tours being purchased.
        case cntOH <= cntSK:
            return nil, float64(cntOH) * -SK.Price
        // Yield a credit for the same number of Sky Tower tours being purchased as the number of
        // Opera House tours being purchased along with extra Sky Tower tours for every Opera House
        // tour being purchased in excess of the number of Sky Tower tours being purchased.
        default:
            var tmp []Tour
            for i := 0; i < (cntOH - cntSK); i++ {
                tmp = append(tmp, SK)
            }
            return tmp, float64(cntSK) * -SK.Price
        }
    }

    // Purchasing more than four (4) Bridge Climb tours yields a discount of $20.00 on all
    // Bridge Climb tours to be purchased.
    SydneyBridgeClimbBulkDiscount := func(t ...Tour) ([]Tour, float64) {
        var cnt int
        for i := range t {
            if t[i].Id == "BC" {
                cnt++
            }
        }
        if cnt > 4 {
            return nil, float64(-20.00 * cnt)
        }
        return nil, 0.00
    }

    tests := [][]Tour{
        []Tour{OH, OH, OH, BC},
        []Tour{OH, SK},
        []Tour{BC, BC, BC, BC, BC, OH},
        []Tour{OH, OH, OH, BC, SK},
        []Tour{OH, BC, BC, SK, SK},
        []Tour{BC, BC, BC, BC, BC, BC, OH, OH},
        []Tour{SK, SK, BC},
    }

    currentPromotions := []SalesPromotion{
        ThreeForTwoOperaHouse,
        FreeSkyTowerWithOperaHouse,
        SydneyBridgeClimbBulkDiscount,
    }

    for t := range tests {
        var sc ShoppingCart
        sc.AddTour(tests[t]...)
        sc.AddSalesPromotion(currentPromotions...)
        sc.Review()
    }

}
