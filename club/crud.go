package club

import (
	"fmt"
	"pkg/utils"
)

var Header = []string{
    "NO",
    "NAMA",
    "PERTANDINGAN",
    "KEMENANGAN",
    "SERI",
    "KEKALAHAN",
    "GOL MEMASUKKAN",
    "GOL KEMASUKKAN",
    "SELISIH GOL",
    "POIN",
}

func CetakRanking(A TabKlub) {
    var i, j int
    var separator string

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
    for i = 0; i < A.N; i++ {
        fmt.Printf(
            "| % 2s | % 4s | % 12d | % 10d | % 4d | % 9d | % 14d | % 14d | % 11d | % 4d |\n",
            utils.ToStr(i+1),
            A.Get[i].Nama,
            A.Get[i].Pertandingan,
            A.Get[i].Menang,
            A.Get[i].Seri,
            A.Get[i].Kalah,
            A.Get[i].Memasukkan,
            A.Get[i].Kemasukkan,
            A.Get[i].Selisih,
            A.Get[i].Poin,
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
