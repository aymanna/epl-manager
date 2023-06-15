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

    C = club.InitialData
    matches.InisialisasiData(&P, C)

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Pilih [1/2/3]: ")
        fmt.Scan(&p)

        if p == "1" {
            club.Menu(&C)
        } else if p == "2" {
            matches.Menu(&P, &C)
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
