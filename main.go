package main

import (
    "fmt"
    "pkg/utils"
    "pkg/club"
    "pkg/matches"
)

var PilihanMenu = []string{
    "Tambah Data Klub Bola",
    "Sunting Data Klub Bola",
    "Hapus Data Klub Bola",
    "Cetak Data Rangkin EPL",
    "Keluar",
}

func main() {
    var p string
    var TKlub = club.DummyData

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Pilih [1/2/3/4/5]: ")
        fmt.Scan(&p)

        if p == "1" {
            club.Menu(&TKlub)
            } else if p == "2" {
            matches.Menu()
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
