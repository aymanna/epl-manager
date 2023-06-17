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
    var p1, p2, nama1, nama2 string
    var week, order, gol1, gol2 int

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Pilih [1/2/3/4/5]: ")
        fmt.Scan(&p1)
        p2 = ""

        if p1 == "1" {          // input
            for p2 != "n" && p2 != "N" {
                fmt.Print("Masukkan nama club (home): ")
                fmt.Scan(&nama1)

                if club.CariKlub(*C, nama1) == -1 {
                    fmt.Println("Klub bola tidak dapat ditemukan.")
                } else {
                    fmt.Print("Masukkan nama club (away): ")
                    fmt.Scan(&nama2)

                    if club.CariKlub(*C, nama2) == -1 {
                        fmt.Println("Klub bola tidak dapat ditemukan.")
                    } else {
                        if nama1 == nama2 {
                            fmt.Println("Club bola home dan away harus berbeda.")
                        } else {
                            week, order = CariPertandingan(*P, nama1, nama2)

                            if P[week][order].SudahMain {
                                fmt.Println("Pertandingan sudah pernah dilakukan.")
                            } else {
                                fmt.Print("Masukkan banyak gol (home): ")
                                fmt.Scan(&gol1)
                                fmt.Print("Masukkan banyak gol (away): ")
                                fmt.Scan(&gol2)

                                if gol1 < 0 || gol2 < 0 {
                                    fmt.Println("Banyak gol tidak dapat bernilai negatif.")
                                } else {
                                    Ubah(P, C, nama1, nama2, gol1, gol2)
                                }
                            }
                        }
                    }
                }

                utils.ValidateRepeat(&p2, "Input data pertandingan lagi? [Y/N]: ")
            }
        } else if p1 == "2" {   // hapus
            for p2 != "n" && p2 != "N" {
                fmt.Print("Masukkan nama club (home): ")
                fmt.Scan(&nama1)

                if club.CariKlub(*C, nama1) == -1 {
                    fmt.Println("Klub bola tidak dapat ditemukan.")
                } else {
                    fmt.Print("Masukkan nama club (away): ")
                    fmt.Scan(&nama2)
    
                    if club.CariKlub(*C, nama2) == -1 {
                        fmt.Println("Klub bola tidak dapat ditemukan.")
                    } else {
                        if nama1 == nama2 {
                            fmt.Println("Club bola home dan away harus berbeda.")
                        } else {
                            Hapus(P, C, nama1, nama2)
                        }
                    }
                }

                utils.ValidateRepeat(&p2, "Hapus data pertandingan lagi? [Y/N]: ")
            }
        } else if p1 == "3" {   // cari
            for p2 != "n" && p2 != "N" {
                club.CetakKlub(*C)
                fmt.Print("Masukkan nama club bola (home): ")
                fmt.Scan(&nama1)

                if club.CariKlub(*C, nama1) == -1 {
                    fmt.Println("Club bola tidak dapat ditemukan.")
                } else {
                    fmt.Print("Masukkan nama club bola (away): ")
                    fmt.Scan(&nama2)

                    if club.CariKlub(*C, nama2) == -1 {
                        fmt.Println("Club bola tidak dapat ditemukan.")
                    } else {
                        if nama1 == nama2 {
                            fmt.Println("Club bola home dan away harus berbeda.")
                        } else {
                            week, order = CariPertandingan(*P, nama1, nama2)
                            fmt.Printf(
                                "Pertandingan %s (home) dengan %s (away) akan berlangsung pada minggu ke-%d dengan urutan pertandingan ke-%d.\n",
                                nama1,
                                nama2,
                                week + 1,
                                order + 1,
                            )
                        }
                    }
                }

                utils.ValidateRepeat(&p2, "Ubah data pertandingan lagi? [Y/N]: ")
            }
        } else if p1 == "4" {   // tampil
            utils.ValidateRepeat(&p2, "Tampil Semuanya? [Y/N]: ")

            if p2 == "Y" || p2 == "y" {
                Cetak(*P)
            } else {
                fmt.Printf("Masukkan minggu pertandingan (1-%d): ", WEEKMAX)
                fmt.Scan(&week)

                if week >= 1 && week <= WEEKMAX {
                    CetakWeek(*P, week-1)
                } else {
                    fmt.Println("Masukkan minggu pertandingan tidak valid.")
                }
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
    fmt.Println(" 4. Tampil Data Pertandingan     ")
    fmt.Println(" 5. Kembali ke Menu Utama        ")
    fmt.Println("---------------------------------")
}
