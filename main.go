package main

import (
	"fmt"
	"pkg/club"
	"pkg/matches"
	"pkg/utils"
)

func main() {
    var p string
    var C club.TabKlub
    var P matches.TabPertandingan

    utils.ValidateRepeat(&p, "Ingin menggunakan nama-nama club tersedia? [Y/N]: ")

    if p == "Y" || p == "y" {
        C = club.InitialData
    }

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Pilih [1/2/3]: ")
        fmt.Scan(&p)

        if p == "1" {
            club.Menu(&C)
        } else if p == "2" {
            if C.N < utils.CLUBMAX {
                fmt.Printf(
                    "Data klub bola masih kurang! (kurang %d lagi)\n",
                    utils.CLUBMAX - C.N,
                )

                fmt.Print(utils.WaitForEnterPrompt)
                utils.WaitForEnter()
            } else {
                matches.InisialisasiData(&P, C)
                matches.Menu(&P, &C)
            }
        } else if p == "3" {
            break
        } else {
            fmt.Println(utils.WrongInputPrompt)
            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
        }
    }

    utils.ClearScreen()
    fmt.Println("Bye")
}

func PrintPrompt() {
    fmt.Println("=================================")
    fmt.Println(" ENGLISH PREMIERE LEAGUE MANAGER ")
    fmt.Println("---------------------------------")
    fmt.Println(" 1. Data Klub Bola               ")
    fmt.Println(" 2. Data Pertandingan            ")
    fmt.Println(" 3. Keluar                       ")
    fmt.Println("---------------------------------")
}
