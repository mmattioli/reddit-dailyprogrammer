//
// Written by Michael Mattioli
//

package main

import (
    "fmt"
    "strings"
    "strconv"
)

var hours = [13]string {
    "twelve",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "ten",
    "eleven",
}

var minutes = map[string][]string {
    "tens" : []string {
        "ten",
        "twenty",
        "thirty",
        "forty",
        "fifty",
    },
    "ones" : hours[1:10],
    "teens" : []string {
        "eleven",
        "twelve",
        "thirteen",
        "fourteen",
        "fifteen",
        "sixteen",
        "seventeen",
        "eighteen",
        "nineteen",
    },
}

// SpokenPhrase takes a string formatted as 24-hour time and returns the corresponding spoken phrase
// used in speech to refer to the particular time of day.
func SpokenPhrase(t string) string {

    h, _ := strconv.Atoi(strings.Split(t, ":")[0])
    m, _ := strconv.Atoi(strings.Split(t, ":")[1])

    hoursWords := func(n int) string {
        return hours[n % 12]
    }

    minutesWords := func(n int) string {
        switch {
        case n == 0:
            return ""
        case n == 10 || n == 20 || n == 30 || n == 40 || n == 50:
            return fmt.Sprintf("%s", minutes["tens"][(n / 10) - 1])
        case n > 0 && n < 10:
            return fmt.Sprintf("oh %s", minutes["ones"][n - 1])
        case n > 10 && n < 20:
            return fmt.Sprintf("%s", minutes["teens"][n - 11])
        default:
            return fmt.Sprintf("%s %s", minutes["tens"][(n / 10) - 1], minutes["ones"][(n % 10) - 1])
        }
    }

    periodWords := func(n int) string {
        if n >= 12 {
            return "pm"
        }
        return "am"
    }

    return fmt.Sprintf("It's %s %s %s", hoursWords(h), minutesWords(m), periodWords(h))
    
}

func main() {
    tests := [7]string{"00:00", "01:30", "12:05", "14:01", "20:29", "21:00", "12:15"}
    for index := range tests {
        fmt.Println(SpokenPhrase(tests[index]))
    }
}
