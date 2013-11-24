package binfo

import (
    "fmt"
    "strconv"
)

func Clump(dna string, k, L, t int) {
    res := make(map[string]int)
    tmp := make(map[string]int)

    for j := 0; j < L; j++ {
        tmp[dna[j:j+k]] = Kmer_occ([]byte(dna[0:L-1]), []byte(dna[j:j+k]))

        if tmp[dna[j:j+k]] >= t {
            res[dna[j:j+k]] = 1
        }
    }

    for inc := 1; inc+L < len(dna); inc++ {
        tmp[dna[inc-1:inc-1+k]] -= 1
        tmp[dna[inc+L-k:inc+L]] += 1

        if tmp[dna[inc+L-k:inc+L]] >= t {
            res[dna[inc+L-k:inc+L]] = 1
        }
    }
    for k, _ := range res {
        fmt.Print(k + " ")
    }

    fmt.Println("\n Number of elements : ", len(res))
}

//FIXME last space must be removed
func Pat_positions(pat, dna string) {
    i := 0

    for inc := 0; inc < len(dna); inc++ {
        for i+inc < len(dna) && dna[inc+i] == pat[i] {
            i++
            if i == len(pat) {
                fmt.Printf(strconv.Itoa(inc) + " ")
                i = 0
                break
            }
        }
        i = 0
    }
}

func Compl_strand(dna string) {
    res := make([]byte, len(dna))

    for inc := 0; inc < len(dna); inc++ {
        if dna[inc] == 'A' {
            res[inc] = 'T'
        } else if dna[inc] == 'T' {
            res[inc] = 'A'
        } else if dna[inc] == 'C' {
            res[inc] = 'G'
        } else if dna[inc] == 'G' {
            res[inc] = 'C'
        }
    }
    fmt.Println(string(res))
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// RETURN the most recurrent pattern of k letters
func Most_kmer(dna string, k int) {
    my := make(map[string]int)
    var temp_wrd string

    for inc := 0; inc < len(dna); inc++ {
        if inc+k < len(dna) {
            temp_wrd = dna[inc : inc+k]
        } else {
            break
        }

        if my[temp_wrd] == 0 {
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

// return the nb of occurrence of a pattern
// We are not using a go function to count overlappings of the pattern
func Kmer_occ(dna, pat []byte) int {
    var nb_occ = 0
    var i = 0

    for inc := 0; inc < len(dna); inc++ {
        for i+inc < len(dna) && dna[inc+i] == pat[i] {
            i++
            if i == len(pat) {
                nb_occ++
                i = 0
                break
            }
        }
        i = 0
    }
    return nb_occ
}
