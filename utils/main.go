package utils

import (
    "fmt"
)

const WrongInputPrompt = "Pilihan anda salah, silahkan input lagi."
const WaitForEnterPrompt = "\nTekan Enter untuk lanjut..."
const CLUBMAX int = 20

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
    var i, last, mid, loops int
    ans := [CLUBMAX/2][2]int{}
    indexes := [CLUBMAX]int{}

    mid = CLUBMAX / 2
    loops = (week - 1) % (CLUBMAX - 1)

    // generate indexes
    for i = 0; i < CLUBMAX; i++ {
        if week <= CLUBMAX - 1 {
            indexes[i] = i
        } else {
            if i < mid {
                indexes[i] = i * 2
            } else {
                indexes[i] = (i - mid) * 2 + 1
            }
        }
    }

    // reorder the indexes based on the match week
    for loops > 0 {
        loops--
        last = indexes[CLUBMAX-1]

        for i = 0; i < CLUBMAX-2; i++ {
            indexes[CLUBMAX-1-i] = indexes[(CLUBMAX-1-i)-1]
        }

        indexes[1] = last
    }

    // insert index of matches
    for i = 0; i < CLUBMAX/2; i++ {
        ans[i] = [2]int{indexes[i], indexes[CLUBMAX-1-i]}
    }

    return ans
}
