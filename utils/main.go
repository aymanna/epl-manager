package utils

import (
    "fmt"
    "strings"
    "os"
    "os/exec"
)

const WrongInputPrompt = "Pilihan anda salah, silahkan input lagi."
const WaitForEnterPrompt = "\nTekan Enter untuk lanjut..."

func ValidateRepeat(p *string, repeatPrompt string) {
    *p = ""

    for *p != "n" && *p != "y" {
        fmt.Print(repeatPrompt)
        fmt.Scan(p)
        *p = strings.TrimSpace(strings.ToLower(*p))
    }
}

func ClearScreen() {
    if isWindows() {
        fmt.Print("\033[H\033[2J")
    } else {
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
    }
}

func isWindows() bool {
    return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}

func WaitForEnter() {
    var s rune
    fmt.Scanf("\n%c", &s)
}
