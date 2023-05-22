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

const KLUBMAX int = 99

type TabKlub struct {
    Get [KLUBMAX]Klub
    N   int
}

var DummyData = TabKlub{
    [KLUBMAX]Klub {
        {"GHI",0,0,0,0,0,0,6,6},
        {"ABC",0,0,0,0,0,0,4,4},
        {"DEF",0,0,0,0,0,0,5,5},
        {"MNO",0,0,0,0,0,0,7,7},
        {"JKL",0,0,0,0,0,0,9,6},
    },
    5,
}

func Menu(A *TabKlub) {
    var p1, p2, nama string
    var k Klub
    var idx int

    for {
        utils.ClearScreen()
        PrintPrompt()
        fmt.Print("Piiih [1/2/3/4/5]: ")
        fmt.Scan(&p1)

        if p1 == "1" {
            if A.N >= KLUBMAX {
                p2 = "n"

                fmt.Println("Data klub bola sudah penuh dan tidak dapat menampung lagi!")
                fmt.Print(utils.WaitForEnterPrompt)
                utils.WaitForEnter()
            }

            for p2 != "n" {
                fmt.Print("Masukkan nama klub: ")
                fmt.Scan(&k.Nama)

                if len(k.Nama) == 3 {
                    fmt.Print("Masukkan banyak pertandingan: ")
                    fmt.Scan(&k.Pertandingan)
                    fmt.Print("Masukkan banyak kemenangan: ")
                    fmt.Scan(&k.Menang)
                    fmt.Print("Masukkan banyak seri: ")
                    fmt.Scan(&k.Seri)
                    fmt.Print("Masukkan banyak kekalahan: ")
                    fmt.Scan(&k.Kalah)
                    fmt.Print("Masukkan jumlah gol memasukkan: ")
                    fmt.Scan(&k.Memasukkan)
                    fmt.Print("Masukkan jumlah gol kemasukkan: ")
                    fmt.Scan(&k.Kemasukkan)

                    k.Selisih = k.Memasukkan - k.Kemasukkan
                    k.Poin = 3 * k.Menang + k.Seri
                    A.Get[A.N] = k
                } else {
                    fmt.Println("panjang nama bola tidak sesuai (harus 3 huruf).")
                }

                if A.N < KLUBMAX {
                    utils.ValidateRepeat(&p2, "Tambahkan data klub bola lagi? [Y/N]: ")
                } else {
                    p2 = "n"
    
                    fmt.Println("Data klub bola sudah penuh")
                    fmt.Print(utils.WaitForEnterPrompt)
                    utils.WaitForEnter()
                }
            }
        } else if p1 == "2" {
            if A.N == 0 {
                p2 = "n"

                fmt.Println("Data klub bola kosong!")
                fmt.Println(utils.WaitForEnterPrompt)
                utils.WaitForEnter()
            } else {
                Cetak(*A)
            }

            for p2 != "n" {
                fmt.Print("Masukkan nama klub bola yang ingin di ubah: ")
                fmt.Scan(&nama)

                idx = CariKlub(*A, nama)

                if idx != -1 {
                    // TODO
                    fmt.Print("Masukkan nama klub baru: ")
                    fmt.Scan(&k.Nama)

                    if len(k.Nama) == 3 {
                        fmt.Print("Masukkan banyak pertandingan baru: ")
                        fmt.Scan(&k.Pertandingan)
                        fmt.Print("Masukkan banyak kemenangan baru: ")
                        fmt.Scan(&k.Menang)
                        fmt.Print("Masukkan banyak seri baru: ")
                        fmt.Scan(&k.Seri)
                        fmt.Print("Masukkan banyak kekalahan baru: ")
                        fmt.Scan(&k.Kalah)
                        fmt.Print("Masukkan jumlah gol memasukkan baru: ")
                        fmt.Scan(&k.Memasukkan)
                        fmt.Print("Masukkan jumlah gol kemasukkan baru: ")
                        fmt.Scan(&k.Kemasukkan)

                        k.Selisih = k.Memasukkan - k.Kemasukkan
                        k.Poin = 3 * k.Menang + k.Seri

                        Ubah(A, nama, k)
                    } else {
                        fmt.Println("panjang nama bola tidak sesuai (harus 3 huruf).")
                    }

                } else {
                    fmt.Println("Nama klub bola tidak dapat ditemukan.")
                }

                utils.ValidateRepeat(&p2, "Ubah data klub bola lagi? [Y/N]: ")
            }
        } else if p1 == "3" {
            if A.N == 0 {
                p2 = "n"

                fmt.Println("Data klub bola kosong!")
                fmt.Println(utils.WaitForEnterPrompt)
                utils.WaitForEnter()
            } else {
                Cetak(*A)
            }

            for p2 != "n" {
                fmt.Print("Masukkan nama klub bola yang ingin di ubah: ")
                fmt.Scan(&nama)

                idx = CariKlub(*A, nama)

                if idx != -1 {
                    Hapus(A, nama)
                } else {
                    fmt.Println("Nama klub bola tidak dapat ditemukan.")
                }

                if A.N > 0 {
                    utils.ValidateRepeat(&p2, "Hapuskan data klub bola lagi? [Y/N]: ")
                } else {
                    p2 = "n"

                    fmt.Println("Data klub bola kosong.")
                    fmt.Print(utils.WaitForEnterPrompt)
                    utils.WaitForEnter()
                }
            }
        } else if p1 == "4" {
            if A.N > 0 {
                Sort(A)
                Cetak(*A)
            } else {
                fmt.Println("Data club bola kosong!")
            }

            fmt.Print(utils.WaitForEnterPrompt)
            utils.WaitForEnter()
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
    fmt.Println("       K L U B     B O L A       ")
    fmt.Println("---------------------------------")
    fmt.Println(" 1. Tambah Data Klub Bola        ")
    fmt.Println(" 2. Ubah Data Klub Bola          ")
    fmt.Println(" 3. Hapus Data Klub Bola         ")
    fmt.Println(" 4. Tampil Data Klub Bola        ")
    fmt.Println(" 5. Kembali ke Menu Utama        ")
    fmt.Println("---------------------------------")
}
