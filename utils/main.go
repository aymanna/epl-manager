package utils

import (
    "fmt"
)

const WrongInputPrompt = "Pilihan anda salah, silahkan input lagi."
const WaitForEnterPrompt = "\nTekan Enter untuk lanjut..."
const CLUBMAX int = 20  // must be in even value

func ValidateRepeat(p *string, repeatPrompt string) {
    *p = ""

    for *p != "n" && *p != "y" && *p != "Y" && *p != "N" {
        fmt.Print(repeatPrompt)
        fmt.Scan(p)
    }
}

func ClearScreen() {
    fmt.Print("\033[H\033[2J")
}

func WaitForEnter() {
    var s rune
    fmt.Scanf("\n%c", &s)
}

func ToStr(n int) string {
    return fmt.Sprintf("%d", n)
}

func DoubleRoundRobin(week int) [CLUBMAX/2][2]int {
    // assume that 1 <= week <= 2 * (CLUBMAX - 1)
    var i, last, loops int
    var ans [CLUBMAX/2][2]int
    var indexes [CLUBMAX]int

    loops = (week - 1) % (CLUBMAX - 1)

    // generate indexes
    for i = 0; i < CLUBMAX; i++ {
        indexes[i] = i
    }

    // reorder the indexes based on the match week
    for loops > 0 {
        loops--
        last = indexes[CLUBMAX-1]

        for i = 0; i < (CLUBMAX-1)-1; i++ {
            indexes[CLUBMAX-1-i] = indexes[(CLUBMAX-1-i)-1]
        }

        indexes[1] = last
    }

    // insert index of matches
    for i = 0; i < CLUBMAX/2; i++ {
        if week <= CLUBMAX - 1 {
            ans[i] = [2]int{indexes[i], indexes[CLUBMAX-1-i]}
        } else {
            ans[i] = [2]int{indexes[CLUBMAX-1-i], indexes[i]}
        }
    }

    return ans
}

func ValidClubName(name string) bool {
    b1 := name[0] >= 'A' && name[0] <= 'Z'
    b2 := name[1] >= 'A' && name[0] <= 'Z'
    b3 := name[2] >= 'A' && name[0] <= 'Z'

    return b1 && b2 && b3
}
