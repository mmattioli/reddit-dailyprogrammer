//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "reflect"
)

// A Tour is an activity that a tourist can partake in.
type Tour struct {
    Id string
    Name string
    Price float64
}

// A SalesPromotion is a rule which determines the eligibility for price adjustments and/or free
// Tours.
type SalesPromotion func(t ...Tour) ([]Tour, float64)

// A ShoppingCart contains the Tours intended for purchase and any SalesPromotions that are to be
// applied to the tourist's order.
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
    fmt.Printf("Items: ")
    var st float64
    for i := range s.Tours {
        st += s.Tours[i].Price
        fmt.Printf("%s ", s.Tours[i].Id)
    }
    fmt.Printf("\nSubtotal: $%.2f\n\n", st)

    fmt.Printf("Sales promotion adjustments\n")
    for i := range s.SalesPromotions {
        tr, adj := s.SalesPromotions[i](s.Tours...)
        if tr != nil {
            fmt.Printf("\tTours added: ")
            for j := range tr {
                fmt.Printf("%s ", tr[j].Id)
            }
        }
        if adj != 0.00 {
            st += adj
            fmt.Printf("\tSubtotal adjustments: $%.2f", adj)
        }
    }

    fmt.Printf("\n\nGrand total: $%.2f\n", st)
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
        return nil, float64(float64(cnt / 3) * -OH.Price)
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
        var tmp []Tour
        if cntOH > 0 {
            for i := 0; i < (cntOH - cntSK); i++ {
                tmp = append(tmp, SK)
            }
        }
        return tmp, float64(cntSK) * -SK.Price
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
            return nil, float64(-20 * cnt)
        }
        return nil, 0.00
    }

    var sc ShoppingCart

    sc.AddTour(OH, OH, OH, BC)
    sc.AddSalesPromotion(ThreeForTwoOperaHouse, FreeSkyTowerWithOperaHouse, SydneyBridgeClimbBulkDiscount)
    sc.Review()

    sc.Empty()

    sc.AddTour(OH, SK)
    sc.AddSalesPromotion(ThreeForTwoOperaHouse, FreeSkyTowerWithOperaHouse, SydneyBridgeClimbBulkDiscount)
    sc.Review()

    sc.Empty()

    sc.AddTour(BC, BC, BC, BC, BC, OH)
    sc.AddSalesPromotion(ThreeForTwoOperaHouse, FreeSkyTowerWithOperaHouse, SydneyBridgeClimbBulkDiscount)
    sc.Review()

}
