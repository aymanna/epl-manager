package matches

import (
    "pkg/club"
    "pkg/utils"
    "fmt"
)

var Header = []string{
    "NO",
    "HOME",
    "AWAY",
    "GOL PERTANDINGAN",
    "SUDAH MAIN",
}

func Cetak(P TabPertandingan, week int) {
    var separator, havePlayed string
    var i, j int

    separator = "+"

    // generate separator for table
    for i = 0; i < len(Header); i++ {
        for j = 0; j < len(Header[i]) + 2; j++ {
            separator += "-"
        }

        separator += "+"
    }

    fmt.Println(separator)

    // print data header
    for i = 0; i < len(Header); i++ {
        fmt.Printf("| %s ", Header[i])
    }

    fmt.Println("|")
    fmt.Println(separator)

    // print contents
    for i = 0; i < PERTANDINGANMAX; i++ {
        p := P[week-1][i]

        if p.SudahMain {
            havePlayed = "SUDAH"
        } else {
            havePlayed = ""    
        }

        fmt.Printf(
            "| % 2s | % 4s | % 4s | % 16s | % 10s |\n",
            utils.ToStr(i+1),
            p.Nama1,
            p.Nama2,
            fmt.Sprintf("%d - %d", p.Gol1, p.Gol2),
            havePlayed,
        )
    }

    fmt.Println(separator)
}

func Ubah(P *TabPertandingan, C *club.TabKlub, week, order, gol1, gol2 int) {
    P[week-1][order-1].Gol1 = gol1
    P[week-1][order-1].Gol2 = gol2
    P[week-1][order-1].SudahMain = true
    UpdateRanking(C, *P)
}

func UbahNama(P *TabPertandingan, namaLama, namaBaru string) {
    for i := 0; i < WEEKMAX; i++ {
        for j := 0; j < PERTANDINGANMAX; j++ {
            if P[i][j].Nama1 == namaLama {
                P[i][j].Nama1 = namaBaru
            } else if P[i][j].Nama2 == namaLama {
                P[i][j].Nama2 = namaBaru
            }
        }
    }
}

func InisialisasiData(P *TabPertandingan, C club.TabKlub) {
    for i := 0; i < WEEKMAX; i++ {
        r := utils.DoubleRoundRobin(i)

        for j := 0; j < PERTANDINGANMAX; j++ {
            P[i][j].Nama1 = C.Get[r[j][0]].Nama
            P[i][j].Nama2 = C.Get[r[j][1]].Nama
            P[i][j].Gol1 = 0
            P[i][j].Gol2 = 0
            P[i][j].SudahMain = false
        }
    }
}

func CariPertandingan(P *TabPertandingan, nama1, nama2 string) (int, int) {
    for i := 0; i < WEEKMAX; i++ {
        for j := 0; j < PERTANDINGANMAX; j++ {
            if P[i][j].Nama1 == nama1 && P[i][j].Nama2 == nama2 {
                return i + 1, j + 1
            }
        }
    }

    return -1, -1
}

func Hapus(P *TabPertandingan, C *club.TabKlub, week, order int) {
    P[week-1][order-1].Gol1 = 0
    P[week-1][order-1].Gol2 = 0
    P[week-1][order-1].SudahMain = false
    UpdateRanking(C, *P)
}

func UpdateRanking(C *club.TabKlub, P TabPertandingan) {
    var i, j, idx1, idx2 int
    var p Pertandingan

    *C = club.InitialData

    for i = 0; i < WEEKMAX; i++ {
        for j = 0; j < PERTANDINGANMAX; j++ {
            p = P[i][j]

            if p.SudahMain {
                idx1 = club.CariKlub(*C, p.Nama1)
                idx2 = club.CariKlub(*C, p.Nama2)

                C.Get[idx1].Pertandingan++
                C.Get[idx2].Pertandingan++

                if p.Gol1 > p.Gol2 {
                    C.Get[idx1].Menang++
                    C.Get[idx2].Kalah++
                } else if p.Gol1 < p.Gol2 {
                    C.Get[idx2].Menang++
                    C.Get[idx1].Kalah++
                } else {
                    C.Get[idx2].Seri++
                    C.Get[idx1].Seri++
                }

                C.Get[idx1].Memasukkan += p.Gol1
                C.Get[idx2].Memasukkan += p.Gol2

                C.Get[idx1].Kemasukkan += p.Gol2
                C.Get[idx2].Kemasukkan += p.Gol1
            }
        }
    }

    C.Get[idx1].Selisih = C.Get[idx1].Memasukkan - C.Get[idx1].Kemasukkan
    C.Get[idx2].Selisih = C.Get[idx2].Memasukkan - C.Get[idx2].Kemasukkan

    C.Get[idx1].Poin = C.Get[idx1].Menang * 3 + C.Get[idx1].Seri
    C.Get[idx2].Poin = C.Get[idx2].Menang * 3 + C.Get[idx2].Seri
}
