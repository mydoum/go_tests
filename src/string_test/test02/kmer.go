package main

import (
    "os"
    "fmt"
    "bytes"
)

var dna []byte
var pattern []byte

func most_kmer(f_dna, f_pat []byte) int {
    var nb_occ = 0
    var i = 0

    for inc := 0; inc < len(f_dna); inc++ {
        for i + inc < len(f_dna) && f_dna[inc + i] == f_pat[i] {
            i++
            if i == len(f_pat) {
                nb_occ++; i = 0
                break
            }
        }
        i = 0
    }
    return nb_occ
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("You must put two args")
        return
    }

    dna = []byte(os.Args[1])
    pattern = []byte(os.Args[2])

    fmt.Println("bytes.Count(dna, Pattern) -> ", bytes.Count(dna, pattern))
    fmt.Println("most_kmer(dna, Pattern) -> ", most_kmer(dna, pattern))
}
