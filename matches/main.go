package matches

import (
	"fmt"
	"pkg/utils"
    "pkg/club"
)

const WEEKMAX int = 2 * (utils.CLUBMAX - 1)
const PERTANDINGANMAX int = (utils.CLUBMAX / 2)

type Pertandingan struct {
    Nama1     string
    Nama2     string
    Gol1      int
    Gol2      int
    SudahMain bool
}

type TabPertandingan [WEEKMAX][PERTANDINGANMAX]Pertandingan

func Menu(P *TabPertandingan, C *club.TabKlub) {
    var p1, p2 string
    var week, order, gol1, gol2 int

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Pilih [1/2/3/4]: ")
        fmt.Scan(&p1)

        if p1 == "1" {          // ubah
            for p2 != "n" && p2 != "N" {
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

                utils.ValidateRepeat(&p2, "Ubah data pertandingan lagi? [Y/N]: ")
            }
        } else if p1 == "2" {   // hapus
            for p2 != "n" && p2 != "N" {
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

                utils.ValidateRepeat(&p2, "Hapus data pertandingan lagi? [Y/N]: ")
            }
        } else if p1 == "3" {   // cari
            // TODO
        } else if p1 == "4" {   // tampil
            fmt.Printf("Masukkan minggu pertandingan (1-%d): ", WEEKMAX)
            fmt.Scan(&week)

            if week >= 1 && week <= WEEKMAX {
                Cetak(*P, week)
            } else {
                fmt.Println("Masukkan minggu pertandingan tidak valid.")
            }

            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
        } else if p1 == "5" {   // exit
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
    fmt.Println(" 1. Input Data Pertandingan      ")
    fmt.Println(" 2. Hapus Data Pertandingan      ")
    fmt.Println(" 3. Cari Data Pertandingan       ")
    fmt.Println(" 3. Tampil Data Pertandingan     ")
    fmt.Println(" 4. Kembali ke Menu Utama        ")
    fmt.Println("---------------------------------")
}
