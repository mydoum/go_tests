package kmer

import "fmt"

func Most_kmer(dna string, k int) {
    my := make(map[string]int)
    var temp_wrd string
    for inc := 0; inc < len(dna); inc++ {
        if inc + k < len(dna) {
            temp_wrd = dna[inc:inc+k]
        } else {
            break
        }

        is_in := false
        for k := range my {
            if k == temp_wrd {
                is_in = true
            }
        }

        if !is_in {
            my[temp_wrd] = Kmer_occ([]byte(dna), []byte(temp_wrd))
        }

    }
    max := 0
    for _, v := range my {
        if v > max {
            max = v
        }
    }

    for k, v := range my {
        if v == max {
            fmt.Print(k + " ")
        }
    }
}

func Kmer_occ(f_dna, f_pat []byte) int {
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
