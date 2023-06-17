package club

import (
	"fmt"
	"pkg/utils"
)

var Header = []string{
    "NO",
    "NAMA",
    "MP",
    "Wins",
    "Draw",
    "Loss",
    "GFs",
    "GAs",
    "GDs",
    "Pts",
}

func CetakRanking(A TabKlub) {
    var i, j int
    var separator, fStr string

    separator = "+"

    // generate separator for table
    for i = 0; i < len(Header); i++ {
        for j = 0; j < len(Header[i]) + 2; j++ {
            separator += "-"
        }

        separator += "+"
    }

    // generate formatted string
    for i = 0; i < len(Header); i++ {
        fStr += fmt.Sprintf("| %% %ds ", len(Header[i]))
    }

    fStr += "|\n"

    fmt.Println(separator)

    // print data header
    for i = 0; i < len(Header); i++ {
        fmt.Printf("| %s ", Header[i])
    }

    fmt.Println("|")
    fmt.Println(separator)

    // print contents
    for i = 0; i < A.N; i++ {
        fmt.Printf(
            fStr,
            utils.ToStr(i+1),
            A.Get[i].Nama,
            utils.ToStr(A.Get[i].Pertandingan),
            utils.ToStr(A.Get[i].Menang),
            utils.ToStr(A.Get[i].Seri),
            utils.ToStr(A.Get[i].Kalah),
            utils.ToStr(A.Get[i].Memasukkan),
            utils.ToStr(A.Get[i].Kemasukkan),
            utils.ToStr(A.Get[i].Selisih),
            utils.ToStr(A.Get[i].Poin),
        )
    }

    fmt.Println(separator)
}

func CetakKlub(A TabKlub) {
    fmt.Println("+------+")
    fmt.Println("| NAMA |")
    fmt.Println("+------+")

    for i := 0; i < A.N; i++ {
        fmt.Printf("| % 4s |\n", A.Get[i].Nama)
    }

    fmt.Println("+------+")
}

func Input(A *TabKlub, nama string) {
    a := Klub{}
    a.Nama = nama
    A.Get[A.N] = a
    A.N++
}

func Ubah(A *TabKlub, namaLama, namaBaru string) {
    A.Get[CariKlub(*A, namaLama)].Nama = namaBaru
}

func Hapus(A *TabKlub, nama string) {
    for i := CariKlub(*A, nama); i < A.N-1; i++ {
        A.Get[i] = A.Get[i+1]
    }

    A.Get[A.N-1] = Klub{}
    A.N--
}

func Sort(A *TabKlub) {
    var i, j int
    var b1, b2 bool

    for i = 1; i < A.N; i++ {
        j = i
        b1 = A.Get[j].Poin > A.Get[j-1].Poin
        b2 = (A.Get[j].Poin == A.Get[j-1].Poin) && (A.Get[j].Selisih > A.Get[j-1].Selisih)

        for j > 0 && (b1 || b2) {
            SwapKlub(&A.Get[j], &A.Get[j-1])
            j--

            if j != 0 {
                b1 = A.Get[j].Poin > A.Get[j-1].Poin
                b2 = (A.Get[j].Poin == A.Get[j-1].Poin) && (A.Get[j].Selisih > A.Get[j-1].Selisih)
            }
        }
    }
}

func SwapKlub(c1, c2 *Klub) {
    tmp := *c1
    *c1 = *c2
    *c2 = tmp
}

func CariKlub(A TabKlub, nama string) int {
    for i := 0; i < A.N; i++ {
        if A.Get[i].Nama == nama {
            return i
        }
    }

    return -1
}

func Drop(A *TabKlub) {
    for i := 0; i < A.N; i++ {
        A.Get[i].Pertandingan = 0
        A.Get[i].Menang = 0
        A.Get[i].Seri = 0
        A.Get[i].Kalah = 0
        A.Get[i].Memasukkan = 0
        A.Get[i].Kemasukkan = 0
        A.Get[i].Selisih = 0
        A.Get[i].Poin = 0
    }
}
