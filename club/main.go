package club

import (
	"fmt"
	"pkg/utils"
)

type Klub struct {
    Nama          string
    Pertandingan  int
    Menang        int
    Seri          int
    Kalah         int
    Memasukkan    int
    Kemasukkan    int
    Selisih       int
    Poin          int
}

type TabKlub struct {
    Get [utils.CLUBMAX]Klub
    N   int
}

var InitialData = TabKlub{
    [utils.CLUBMAX]Klub {
        {"MCI",0,0,0,0,0,0,0,0},
        {"ARS",0,0,0,0,0,0,0,0},
        {"MUN",0,0,0,0,0,0,0,0},
        {"NEW",0,0,0,0,0,0,0,0},
        {"LIV",0,0,0,0,0,0,0,0},
        {"BHA",0,0,0,0,0,0,0,0},
        {"AVL",0,0,0,0,0,0,0,0},
        {"TOT",0,0,0,0,0,0,0,0},
        {"BRE",0,0,0,0,0,0,0,0},
        {"FUL",0,0,0,0,0,0,0,0},
        {"CRY",0,0,0,0,0,0,0,0},
        {"CHE",0,0,0,0,0,0,0,0},
        {"WOL",0,0,0,0,0,0,0,0},
        {"WHU",0,0,0,0,0,0,0,0},
        {"BOU",0,0,0,0,0,0,0,0},
        {"NFO",0,0,0,0,0,0,0,0},
        {"EVE",0,0,0,0,0,0,0,0},
        {"LEI",0,0,0,0,0,0,0,0},
        {"LEE",0,0,0,0,0,0,0,0},
        {"SOU",0,0,0,0,0,0,0,0},
    },
    20,
}

func Menu(A *TabKlub) {
    var p1, p2, nama, namaBaru string
    var idx int

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Piiih [1/2/3]: ")
        fmt.Scan(&p1)

        if p1 == "1" {   // ubah
            if A.N == 0 {
                p2 = "n"

                fmt.Println("Data klub bola kosong!")
                fmt.Print(utils.WaitForEnterPrompt)
                utils.WaitForEnter()
            } else {
                Cetak(*A)
            }

            for p2 != "n" && p2 != "N" {
                fmt.Print("Masukkan nama klub bola yang ingin di ubah: ")
                fmt.Scan(&nama)

                idx = CariKlub(*A, nama)

                if idx != -1 {
                    fmt.Print("Masukkan nama klub baru: ")
                    fmt.Scan(&namaBaru)

                    if len(namaBaru) == 3 {
                        Ubah(A, nama, namaBaru)
                    } else {
                        fmt.Println("panjang nama bola tidak sesuai (harus 3 huruf).")
                    }

                } else {
                    fmt.Println("Nama klub bola tidak dapat ditemukan.")
                }

                utils.ValidateRepeat(&p2, "Ubah data klub bola lagi? [Y/N]: ")
            }
        } else if p1 == "2" {   // tampil
            if A.N > 0 {
                // UpdateRanking() ???
                Sort(A)
                Cetak(*A)
            } else {
                fmt.Println("Data club bola kosong!")
            }

            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
        } else if p1 == "3" {   // exit
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
    fmt.Println("       K L U B     B O L A       ")
    fmt.Println("---------------------------------")
    fmt.Println(" 1. Ubah Data Klub Bola          ")
    fmt.Println(" 2. Tampil Ranking Klub Bola     ")
    fmt.Println(" 3. Kembali ke Menu Utama        ")
    fmt.Println("---------------------------------")
}
