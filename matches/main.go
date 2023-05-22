package matches

import (
	"fmt"
	"pkg/utils"
)

type Pertandingan struct {
    Club1 string
    Club2 string
    Gol1  int
    Gol2  int
}

const PERTANDINGANMAX int = 1000

var DummyData = [PERTANDINGANMAX]Pertandingan{}

func Menu() {
    var p1 string

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Pilih [1/2/3/4/5]: ")
        fmt.Scan(&p1)

        if p1 == "1" {
            // TODO
        } else if p1 == "2" {
            // TODO
        } else if p1 == "3" {
            // TODO
        } else if p1 == "4" {
            // TODO
        } else if p1 == "5" {
            break
        } else {
            fmt.Println(utils.WrongInputPrompt)
            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
        }
    }
}

func PrintPrompt() {
    fmt.Println("=================================")
    fmt.Println("     P E R T A N D I N G A N     ")
    fmt.Println("---------------------------------")
    fmt.Println(" 1. Tambah Data Pertandingan     ")
    fmt.Println(" 2. Ubah Data Pertandingan       ")
    fmt.Println(" 3. Hapus Data Pertandingan      ")
    fmt.Println(" 4. Tampil Data Pertandingan     ")
    fmt.Println(" 5. Kembali ke Menu Utama        ")
    fmt.Println("---------------------------------")
}
