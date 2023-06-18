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

var DummyData = TabKlub{
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
    var A_tmp TabKlub

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Piiih [1/2/3/4/5]: ")
        fmt.Scan(&p1)
        p2 = ""

        if p1 == "1" {          // input
            if A.N >= utils.CLUBMAX {
                fmt.Println("Data club bola sudah penuh.")
                fmt.Print(utils.WaitForEnterPrompt)
                utils.WaitForEnter()

                p2 = "N"
            }

            for p2 != "n" && p2 != "N" {
                fmt.Print("Masukkan nama club bola yang ingin di masukkan (3 alfabetikal caps): ")
                fmt.Scan(&nama)

                if CariKlub(*A, nama) != -1 {
                    fmt.Println("Nama club bola sudah ada.")
                } else {
                    if len(nama) != 3 {
                        fmt.Println("panjang nama bola tidak sesuai (harus 3 huruf).")
                    } else {
                        if !utils.ValidClubName(nama) {
                            fmt.Println("Nama klub tidak valid (diperlukan semua caps dan alfabetikal).")
                        } else {
                            Input(A, nama)
                        }
                    }
                }

                utils.ValidateRepeat(&p2, "Input data klub bola lagi? [Y/N]: ")
            }
        } else if p1 == "2" {   // ubah
            if A.N == 0 {
                p2 = "n"

                fmt.Println("Data klub bola kosong!")
                fmt.Print(utils.WaitForEnterPrompt)
                utils.WaitForEnter()
            } else {
                CetakKlub(*A)
            }

            for p2 != "n" && p2 != "N" {
                fmt.Print("Masukkan nama klub bola yang ingin di ubah: ")
                fmt.Scan(&nama)

                if CariKlub(*A, nama) == -1 {
                    fmt.Println("Nama klub bola tidak dapat ditemukan.")
                } else {
                    fmt.Print("Masukkan nama klub baru: ")
                    fmt.Scan(&namaBaru)

                    if len(namaBaru) == 3 {
                        if CariKlub(*A, namaBaru) != -1 {
                            fmt.Println("Nama klub bola sudah ada.")
                        }

                        if !utils.ValidClubName(namaBaru) {
                            fmt.Println("Nama klub tidak valid (diperlukan semua caps dan alfabetikal).")
                        } else {
                            Ubah(A, nama, namaBaru)
                        }
                    } else {
                        fmt.Println("panjang nama bola tidak sesuai (harus 3 huruf).")
                    }
                }

                utils.ValidateRepeat(&p2, "Ubah data klub bola lagi? [Y/N]: ")
            }
        } else if p1 == "3" {   // hapus
            if !(A.N > 0) {
                fmt.Println("Data club bola masih kosong!")
                p2 = "N"
            }

            for p2 != "n" && p2 != "N" {
                CetakKlub(*A)

                fmt.Print("Masukkan nama club bola yang ingin dihapus: ")
                fmt.Scan(&nama)

                if CariKlub(*A, nama) == -1 {
                    fmt.Println("Nama club bola tidak dapat ditemukan.")
                } else {
                    Hapus(A, nama)
                }

                utils.ValidateRepeat(&p2, "Hapus data klub bola lagi? [Y/N]: ")
            }
        } else if p1 == "4" {   // tampil
            if A.N > 0 {
                A_tmp = *A
                utils.ValidateRepeat(&p2, "Data terurut secara descending (menurun)? [Y/N]: ")

                if p2 == "Y" || p2 == "y" {
                    InsertionSort(&A_tmp)
                } else {
                    SelectionSort(&A_tmp)
                }

                CetakRanking(A_tmp)
            } else {
                fmt.Println("Data club bola kosong!")
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
    fmt.Println("       K L U B     B O L A       ")
    fmt.Println("---------------------------------")
    fmt.Println(" 1. Input Data Klub Bola         ")
    fmt.Println(" 2. Ubah Data Klub Bola          ")
    fmt.Println(" 3. Hapus Data Klub Bola         ")
    fmt.Println(" 4. Tampil Ranking Klub Bola     ")
    fmt.Println(" 5. Kembali ke Menu Utama        ")
    fmt.Println("---------------------------------")
}
