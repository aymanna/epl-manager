package club

import (
    "fmt"
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

func Cetak(A TabKlub) {
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
            "| % d | % 4s | % 12d | % 10d | % 4d | % 9d | % 14d | % 14d | % 11d | % 4d |\n",
            i+1,
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

func Ubah(A *TabKlub, nama string, k Klub) {
    A.Get[CariKlub(*A, nama)] = k
}

func Hapus(A *TabKlub, nama string) {
    for i := CariKlub(*A, nama); i < A.N-1; i++ {
        A.Get[i] = A.Get[i+1]
    }

    A.Get[A.N] = Klub{}
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
            Swap(&A.Get[j], &A.Get[j-1])
            j--

            if j != 0 {
                b1 = A.Get[j].Poin > A.Get[j-1].Poin
                b2 = (A.Get[j].Poin == A.Get[j-1].Poin) && (A.Get[j].Selisih > A.Get[j-1].Selisih)
            }
        }
    }
}

func Swap(c1, c2 *Klub) {
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
