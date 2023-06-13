package matches

import (
	"fmt"
	"pkg/utils"
    "pkg/club"
)

const WEEKMAX int = 2 * (utils.CLUBMAX - 1)
const PERTANDINGANMAX int = (utils.CLUBMAX / 2)

type Pertandingan struct {
    Nama1      string
    Nama2      string
    Gol1       int
    Gol2       int
    HavePlayed bool
}

type TabPertandingan [WEEKMAX][PERTANDINGANMAX]Pertandingan

var InitialData = TabPertandingan{
    {
        {"MCI", "SOU", 0, 0, false},
        {"ARS", "LEE", 0, 0, false},
        {"MUN", "LEI", 0, 0, false},
        {"NEW", "EVE", 0, 0, false},
        {"LIV", "NFO", 0, 0, false},
        {"BHA", "BOU", 0, 0, false},
        {"AVL", "WHU", 0, 0, false},
        {"TOT", "WOL", 0, 0, false},
        {"BRE", "CHE", 0, 0, false},
        {"FUL", "CRY", 0, 0, false},
    },
    {
        {"MCI", "LEE", 0, 0, false},
        {"SOU", "LEI", 0, 0, false},
        {"ARS", "EVE", 0, 0, false},
        {"MUN", "NFO", 0, 0, false},
        {"NEW", "BOU", 0, 0, false},
        {"LIV", "WHU", 0, 0, false},
        {"BHA", "WOL", 0, 0, false},
        {"AVL", "CHE", 0, 0, false},
        {"TOT", "CRY", 0, 0, false},
        {"BRE", "FUL", 0, 0, false},
    },
    {
        {"MCI", "LEI", 0, 0, false},
        {"LEE", "EVE", 0, 0, false},
        {"SOU", "NFO", 0, 0, false},
        {"ARS", "BOU", 0, 0, false},
        {"MUN", "WHU", 0, 0, false},
        {"NEW", "WOL", 0, 0, false},
        {"LIV", "CHE", 0, 0, false},
        {"BHA", "CRY", 0, 0, false},
        {"AVL", "FUL", 0, 0, false},
        {"TOT", "BRE", 0, 0, false},
    },
    {
        {"MCI", "EVE", 0, 0, false},
        {"LEI", "NFO", 0, 0, false},
        {"LEE", "BOU", 0, 0, false},
        {"SOU", "WHU", 0, 0, false},
        {"ARS", "WOL", 0, 0, false},
        {"MUN", "CHE", 0, 0, false},
        {"NEW", "CRY", 0, 0, false},
        {"LIV", "FUL", 0, 0, false},
        {"BHA", "BRE", 0, 0, false},
        {"AVL", "TOT", 0, 0, false},
    },
    {
        {"MCI", "NFO", 0, 0, false},
        {"EVE", "BOU", 0, 0, false},
        {"LEI", "WHU", 0, 0, false},
        {"LEE", "WOL", 0, 0, false},
        {"SOU", "CHE", 0, 0, false},
        {"ARS", "CRY", 0, 0, false},
        {"MUN", "FUL", 0, 0, false},
        {"NEW", "BRE", 0, 0, false},
        {"LIV", "TOT", 0, 0, false},
        {"BHA", "AVL", 0, 0, false},
    },
    {
        {"MCI", "BOU", 0, 0, false},
        {"NFO", "WHU", 0, 0, false},
        {"EVE", "WOL", 0, 0, false},
        {"LEI", "CHE", 0, 0, false},
        {"LEE", "CRY", 0, 0, false},
        {"SOU", "FUL", 0, 0, false},
        {"ARS", "BRE", 0, 0, false},
        {"MUN", "TOT", 0, 0, false},
        {"NEW", "AVL", 0, 0, false},
        {"LIV", "BHA", 0, 0, false},
    },
    {
        {"MCI", "WHU", 0, 0, false},
        {"BOU", "WOL", 0, 0, false},
        {"NFO", "CHE", 0, 0, false},
        {"EVE", "CRY", 0, 0, false},
        {"LEI", "FUL", 0, 0, false},
        {"LEE", "BRE", 0, 0, false},
        {"SOU", "TOT", 0, 0, false},
        {"ARS", "AVL", 0, 0, false},
        {"MUN", "BHA", 0, 0, false},
        {"NEW", "LIV", 0, 0, false},
    },
    {
        {"MCI", "WOL", 0, 0, false},
        {"WHU", "CHE", 0, 0, false},
        {"BOU", "CRY", 0, 0, false},
        {"NFO", "FUL", 0, 0, false},
        {"EVE", "BRE", 0, 0, false},
        {"LEI", "TOT", 0, 0, false},
        {"LEE", "AVL", 0, 0, false},
        {"SOU", "BHA", 0, 0, false},
        {"ARS", "LIV", 0, 0, false},
        {"MUN", "NEW", 0, 0, false},
    },
    {
        {"MCI", "CHE", 0, 0, false},
        {"WOL", "CRY", 0, 0, false},
        {"WHU", "FUL", 0, 0, false},
        {"BOU", "BRE", 0, 0, false},
        {"NFO", "TOT", 0, 0, false},
        {"EVE", "AVL", 0, 0, false},
        {"LEI", "BHA", 0, 0, false},
        {"LEE", "LIV", 0, 0, false},
        {"SOU", "NEW", 0, 0, false},
        {"ARS", "MUN", 0, 0, false},
    },
    {
        {"MCI", "CRY", 0, 0, false},
        {"CHE", "FUL", 0, 0, false},
        {"WOL", "BRE", 0, 0, false},
        {"WHU", "TOT", 0, 0, false},
        {"BOU", "AVL", 0, 0, false},
        {"NFO", "BHA", 0, 0, false},
        {"EVE", "LIV", 0, 0, false},
        {"LEI", "NEW", 0, 0, false},
        {"LEE", "MUN", 0, 0, false},
        {"SOU", "ARS", 0, 0, false},
    },
    {
        {"MCI", "FUL", 0, 0, false},
        {"CRY", "BRE", 0, 0, false},
        {"CHE", "TOT", 0, 0, false},
        {"WOL", "AVL", 0, 0, false},
        {"WHU", "BHA", 0, 0, false},
        {"BOU", "LIV", 0, 0, false},
        {"NFO", "NEW", 0, 0, false},
        {"EVE", "MUN", 0, 0, false},
        {"LEI", "ARS", 0, 0, false},
        {"LEE", "SOU", 0, 0, false},
    },
    {
        {"MCI", "BRE", 0, 0, false},
        {"FUL", "TOT", 0, 0, false},
        {"CRY", "AVL", 0, 0, false},
        {"CHE", "BHA", 0, 0, false},
        {"WOL", "LIV", 0, 0, false},
        {"WHU", "NEW", 0, 0, false},
        {"BOU", "MUN", 0, 0, false},
        {"NFO", "ARS", 0, 0, false},
        {"EVE", "SOU", 0, 0, false},
        {"LEI", "LEE", 0, 0, false},
    },
    {
        {"MCI", "TOT", 0, 0, false},
        {"BRE", "AVL", 0, 0, false},
        {"FUL", "BHA", 0, 0, false},
        {"CRY", "LIV", 0, 0, false},
        {"CHE", "NEW", 0, 0, false},
        {"WOL", "MUN", 0, 0, false},
        {"WHU", "ARS", 0, 0, false},
        {"BOU", "SOU", 0, 0, false},
        {"NFO", "LEE", 0, 0, false},
        {"EVE", "LEI", 0, 0, false},
    },
    {
        {"MCI", "AVL", 0, 0, false},
        {"TOT", "BHA", 0, 0, false},
        {"BRE", "LIV", 0, 0, false},
        {"FUL", "NEW", 0, 0, false},
        {"CRY", "MUN", 0, 0, false},
        {"CHE", "ARS", 0, 0, false},
        {"WOL", "SOU", 0, 0, false},
        {"WHU", "LEE", 0, 0, false},
        {"BOU", "LEI", 0, 0, false},
        {"NFO", "EVE", 0, 0, false},
    },
    {
        {"MCI", "BHA", 0, 0, false},
        {"AVL", "LIV", 0, 0, false},
        {"TOT", "NEW", 0, 0, false},
        {"BRE", "MUN", 0, 0, false},
        {"FUL", "ARS", 0, 0, false},
        {"CRY", "SOU", 0, 0, false},
        {"CHE", "LEE", 0, 0, false},
        {"WOL", "LEI", 0, 0, false},
        {"WHU", "EVE", 0, 0, false},
        {"BOU", "NFO", 0, 0, false},
    },
    {
        {"MCI", "LIV", 0, 0, false},
        {"BHA", "NEW", 0, 0, false},
        {"AVL", "MUN", 0, 0, false},
        {"TOT", "ARS", 0, 0, false},
        {"BRE", "SOU", 0, 0, false},
        {"FUL", "LEE", 0, 0, false},
        {"CRY", "LEI", 0, 0, false},
        {"CHE", "EVE", 0, 0, false},
        {"WOL", "NFO", 0, 0, false},
        {"WHU", "BOU", 0, 0, false},
    },
    {
        {"MCI", "NEW", 0, 0, false},
        {"LIV", "MUN", 0, 0, false},
        {"BHA", "ARS", 0, 0, false},
        {"AVL", "SOU", 0, 0, false},
        {"TOT", "LEE", 0, 0, false},
        {"BRE", "LEI", 0, 0, false},
        {"FUL", "EVE", 0, 0, false},
        {"CRY", "NFO", 0, 0, false},
        {"CHE", "BOU", 0, 0, false},
        {"WOL", "WHU", 0, 0, false},
    },
    {
        {"MCI", "MUN", 0, 0, false},
        {"NEW", "ARS", 0, 0, false},
        {"LIV", "SOU", 0, 0, false},
        {"BHA", "LEE", 0, 0, false},
        {"AVL", "LEI", 0, 0, false},
        {"TOT", "EVE", 0, 0, false},
        {"BRE", "NFO", 0, 0, false},
        {"FUL", "BOU", 0, 0, false},
        {"CRY", "WHU", 0, 0, false},
        {"CHE", "WOL", 0, 0, false},
    },
    {
        {"MCI", "ARS", 0, 0, false},
        {"MUN", "SOU", 0, 0, false},
        {"NEW", "LEE", 0, 0, false},
        {"LIV", "LEI", 0, 0, false},
        {"BHA", "EVE", 0, 0, false},
        {"AVL", "NFO", 0, 0, false},
        {"TOT", "BOU", 0, 0, false},
        {"BRE", "WHU", 0, 0, false},
        {"FUL", "WOL", 0, 0, false},
        {"CRY", "CHE", 0, 0, false},
    },
    {
        {"MCI", "SOU", 0, 0, false},
        {"MUN", "LEI", 0, 0, false},
        {"LIV", "NFO", 0, 0, false},
        {"AVL", "WHU", 0, 0, false},
        {"BRE", "CHE", 0, 0, false},
        {"CRY", "FUL", 0, 0, false},
        {"WOL", "TOT", 0, 0, false},
        {"BOU", "BHA", 0, 0, false},
        {"EVE", "NEW", 0, 0, false},
        {"LEE", "ARS", 0, 0, false},
    },
    {
        {"MCI", "LEI", 0, 0, false},
        {"SOU", "NFO", 0, 0, false},
        {"MUN", "WHU", 0, 0, false},
        {"LIV", "CHE", 0, 0, false},
        {"AVL", "FUL", 0, 0, false},
        {"BRE", "TOT", 0, 0, false},
        {"CRY", "BHA", 0, 0, false},
        {"WOL", "NEW", 0, 0, false},
        {"BOU", "ARS", 0, 0, false},
        {"EVE", "LEE", 0, 0, false},
    },
    {
        {"MCI", "NFO", 0, 0, false},
        {"LEI", "WHU", 0, 0, false},
        {"SOU", "CHE", 0, 0, false},
        {"MUN", "FUL", 0, 0, false},
        {"LIV", "TOT", 0, 0, false},
        {"AVL", "BHA", 0, 0, false},
        {"BRE", "NEW", 0, 0, false},
        {"CRY", "ARS", 0, 0, false},
        {"WOL", "LEE", 0, 0, false},
        {"BOU", "EVE", 0, 0, false},
    },
    {
        {"MCI", "WHU", 0, 0, false},
        {"NFO", "CHE", 0, 0, false},
        {"LEI", "FUL", 0, 0, false},
        {"SOU", "TOT", 0, 0, false},
        {"MUN", "BHA", 0, 0, false},
        {"LIV", "NEW", 0, 0, false},
        {"AVL", "ARS", 0, 0, false},
        {"BRE", "LEE", 0, 0, false},
        {"CRY", "EVE", 0, 0, false},
        {"WOL", "BOU", 0, 0, false},
    },
    {
        {"MCI", "CHE", 0, 0, false},
        {"WHU", "FUL", 0, 0, false},
        {"NFO", "TOT", 0, 0, false},
        {"LEI", "BHA", 0, 0, false},
        {"SOU", "NEW", 0, 0, false},
        {"MUN", "ARS", 0, 0, false},
        {"LIV", "LEE", 0, 0, false},
        {"AVL", "EVE", 0, 0, false},
        {"BRE", "BOU", 0, 0, false},
        {"CRY", "WOL", 0, 0, false},
    },
    {
        {"MCI", "FUL", 0, 0, false},
        {"CHE", "TOT", 0, 0, false},
        {"WHU", "BHA", 0, 0, false},
        {"NFO", "NEW", 0, 0, false},
        {"LEI", "ARS", 0, 0, false},
        {"SOU", "LEE", 0, 0, false},
        {"MUN", "EVE", 0, 0, false},
        {"LIV", "BOU", 0, 0, false},
        {"AVL", "WOL", 0, 0, false},
        {"BRE", "CRY", 0, 0, false},
    },
    {
        {"MCI", "TOT", 0, 0, false},
        {"FUL", "BHA", 0, 0, false},
        {"CHE", "NEW", 0, 0, false},
        {"WHU", "ARS", 0, 0, false},
        {"NFO", "LEE", 0, 0, false},
        {"LEI", "EVE", 0, 0, false},
        {"SOU", "BOU", 0, 0, false},
        {"MUN", "WOL", 0, 0, false},
        {"LIV", "CRY", 0, 0, false},
        {"AVL", "BRE", 0, 0, false},
    },
    {
        {"MCI", "BHA", 0, 0, false},
        {"TOT", "NEW", 0, 0, false},
        {"FUL", "ARS", 0, 0, false},
        {"CHE", "LEE", 0, 0, false},
        {"WHU", "EVE", 0, 0, false},
        {"NFO", "BOU", 0, 0, false},
        {"LEI", "WOL", 0, 0, false},
        {"SOU", "CRY", 0, 0, false},
        {"MUN", "BRE", 0, 0, false},
        {"LIV", "AVL", 0, 0, false},
    },
    {
        {"MCI", "NEW", 0, 0, false},
        {"BHA", "ARS", 0, 0, false},
        {"TOT", "LEE", 0, 0, false},
        {"FUL", "EVE", 0, 0, false},
        {"CHE", "BOU", 0, 0, false},
        {"WHU", "WOL", 0, 0, false},
        {"NFO", "CRY", 0, 0, false},
        {"LEI", "BRE", 0, 0, false},
        {"SOU", "AVL", 0, 0, false},
        {"MUN", "LIV", 0, 0, false},
    },
    {
        {"MCI", "ARS", 0, 0, false},
        {"NEW", "LEE", 0, 0, false},
        {"BHA", "EVE", 0, 0, false},
        {"TOT", "BOU", 0, 0, false},
        {"FUL", "WOL", 0, 0, false},
        {"CHE", "CRY", 0, 0, false},
        {"WHU", "BRE", 0, 0, false},
        {"NFO", "AVL", 0, 0, false},
        {"LEI", "LIV", 0, 0, false},
        {"SOU", "MUN", 0, 0, false},
    },
    {
        {"MCI", "LEE", 0, 0, false},
        {"ARS", "EVE", 0, 0, false},
        {"NEW", "BOU", 0, 0, false},
        {"BHA", "WOL", 0, 0, false},
        {"TOT", "CRY", 0, 0, false},
        {"FUL", "BRE", 0, 0, false},
        {"CHE", "AVL", 0, 0, false},
        {"WHU", "LIV", 0, 0, false},
        {"NFO", "MUN", 0, 0, false},
        {"LEI", "SOU", 0, 0, false},
    },
    {
        {"MCI", "EVE", 0, 0, false},
        {"LEE", "BOU", 0, 0, false},
        {"ARS", "WOL", 0, 0, false},
        {"NEW", "CRY", 0, 0, false},
        {"BHA", "BRE", 0, 0, false},
        {"TOT", "AVL", 0, 0, false},
        {"FUL", "LIV", 0, 0, false},
        {"CHE", "MUN", 0, 0, false},
        {"WHU", "SOU", 0, 0, false},
        {"NFO", "LEI", 0, 0, false},
    },
    {
        {"MCI", "BOU", 0, 0, false},
        {"EVE", "WOL", 0, 0, false},
        {"LEE", "CRY", 0, 0, false},
        {"ARS", "BRE", 0, 0, false},
        {"NEW", "AVL", 0, 0, false},
        {"BHA", "LIV", 0, 0, false},
        {"TOT", "MUN", 0, 0, false},
        {"FUL", "SOU", 0, 0, false},
        {"CHE", "LEI", 0, 0, false},
        {"WHU", "NFO", 0, 0, false},
    },
    {
        {"MCI", "WOL", 0, 0, false},
        {"BOU", "CRY", 0, 0, false},
        {"EVE", "BRE", 0, 0, false},
        {"LEE", "AVL", 0, 0, false},
        {"ARS", "LIV", 0, 0, false},
        {"NEW", "MUN", 0, 0, false},
        {"BHA", "SOU", 0, 0, false},
        {"TOT", "LEI", 0, 0, false},
        {"FUL", "NFO", 0, 0, false},
        {"CHE", "WHU", 0, 0, false},
    },
    {
        {"MCI", "CRY", 0, 0, false},
        {"WOL", "BRE", 0, 0, false},
        {"BOU", "AVL", 0, 0, false},
        {"EVE", "LIV", 0, 0, false},
        {"LEE", "MUN", 0, 0, false},
        {"ARS", "SOU", 0, 0, false},
        {"NEW", "LEI", 0, 0, false},
        {"BHA", "NFO", 0, 0, false},
        {"TOT", "WHU", 0, 0, false},
        {"FUL", "CHE", 0, 0, false},
    },
    {
        {"MCI", "BRE", 0, 0, false},
        {"CRY", "AVL", 0, 0, false},
        {"WOL", "LIV", 0, 0, false},
        {"BOU", "MUN", 0, 0, false},
        {"EVE", "SOU", 0, 0, false},
        {"LEE", "LEI", 0, 0, false},
        {"ARS", "NFO", 0, 0, false},
        {"NEW", "WHU", 0, 0, false},
        {"BHA", "CHE", 0, 0, false},
        {"TOT", "FUL", 0, 0, false},
    },
    {
        {"MCI", "AVL", 0, 0, false},
        {"BRE", "LIV", 0, 0, false},
        {"CRY", "MUN", 0, 0, false},
        {"WOL", "SOU", 0, 0, false},
        {"BOU", "LEI", 0, 0, false},
        {"EVE", "NFO", 0, 0, false},
        {"LEE", "WHU", 0, 0, false},
        {"ARS", "CHE", 0, 0, false},
        {"NEW", "FUL", 0, 0, false},
        {"BHA", "TOT", 0, 0, false},
    },
    {
        {"MCI", "LIV", 0, 0, false},
        {"AVL", "MUN", 0, 0, false},
        {"BRE", "SOU", 0, 0, false},
        {"CRY", "LEI", 0, 0, false},
        {"WOL", "NFO", 0, 0, false},
        {"BOU", "WHU", 0, 0, false},
        {"EVE", "CHE", 0, 0, false},
        {"LEE", "FUL", 0, 0, false},
        {"ARS", "TOT", 0, 0, false},
        {"NEW", "BHA", 0, 0, false},
    },
    {
        {"MCI", "MUN", 0, 0, false},
        {"LIV", "SOU", 0, 0, false},
        {"AVL", "LEI", 0, 0, false},
        {"BRE", "NFO", 0, 0, false},
        {"CRY", "WHU", 0, 0, false},
        {"WOL", "CHE", 0, 0, false},
        {"BOU", "FUL", 0, 0, false},
        {"EVE", "TOT", 0, 0, false},
        {"LEE", "BHA", 0, 0, false},
        {"ARS", "NEW", 0, 0, false},
    },       
}

func Menu(P *TabPertandingan, C *club.TabKlub) {
    var p1 string
    var week, order, gol1, gol2 int

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Pilih [1/2/3/4]: ")
        fmt.Scan(&p1)

        if p1 == "1" {          // ubah
            fmt.Printf("Masukkan minggu pertandingan (1-%d): ", WEEKMAX)
            fmt.Scan(&week)

            if week >= 1 && week <= WEEKMAX {
                Cetak(*P, week)
                fmt.Println()

                fmt.Print("Masukkan nomor pertandingan: ")
                fmt.Scan(&order)

                if order < 1 && order > PERTANDINGANMAX {
                    fmt.Println("Masukkan minggu pertandingan tidak valid.")
                } else {
                    fmt.Print("Masukkan banyak gol (home): ")
                    fmt.Scan(&gol1)
                    fmt.Print("Masukkan banyak gol (away): ")
                    fmt.Scan(&gol2)

                    if gol1 < 0 || gol2 < 0 {
                        fmt.Println("Banyak gol tidak dapat bernilai negatif.")
                    } else {
                        Ubah(P, C, week, order, gol1, gol2)
                    }
                }
            } else {
                fmt.Println("Masukkan minggu pertandingan tidak valid.")
            }

            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
        } else if p1 == "2" {   // hapus
            fmt.Printf("Masukkan minggu pertandingan (1-%d): ", WEEKMAX)
            fmt.Scan(&week)

            if week > 0 && week < WEEKMAX {
                Cetak(*P, week)
                fmt.Println()

                fmt.Print("Masukkan nomor pertandingan: ")
                fmt.Scan(&order)

                if order < 1 && order > PERTANDINGANMAX {
                    fmt.Println("Masukkan minggu pertandingan tidak valid.")
                } else {
                    Hapus(P, C, week, order)
                }
            } else {
                fmt.Println("Masukkan minggu pertandingan tidak valid.")
            }

            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
        } else if p1 == "3" {   // tampil
            fmt.Printf("Masukkan minggu pertandingan (1-%d): ", WEEKMAX)
            fmt.Scan(&week)

            if week >= 1 && week <= WEEKMAX {
                Cetak(*P, week)
            } else {
                fmt.Println("Masukkan minggu pertandingan tidak valid.")
            }

            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
        } else if p1 == "4" {   // exit
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
    fmt.Println(" 1. Ubah Data Pertandingan       ")
    fmt.Println(" 2. Hapus Data Pertandingan      ")
    fmt.Println(" 3. Tampil Data Pertandingan     ")
    fmt.Println(" 4. Kembali ke Menu Utama        ")
    fmt.Println("---------------------------------")
}
