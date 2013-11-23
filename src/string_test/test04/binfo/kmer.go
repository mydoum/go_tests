package binfo

import (
    "container/list"
    "fmt"
    "strconv"
)

// FIXME try to use dictionnary instead of a list
func Clump(dna string, k, L, t int) {
    res := list.New()
    tmp := list.New()

    for inc := 0; inc < len(dna); inc++ {
        if inc+L >= len(dna) {
            break
        } else {
            sub := dna[inc : inc+L]
            tmp = list_kmer(sub, k, t)
            is_pres := false

            for e := tmp.Front(); e != nil; e = e.Next() {
                for f := res.Front(); f != nil; f = f.Next() {
                    if f.Value == e.Value {
                        is_pres = true
                        break
                    }
                }
                if !is_pres {
                    res.PushBack(e.Value)
                }
                is_pres = false
            }
        }
    }
    for e := res.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
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

// FIXME optimize it using a mapping
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

// FIXME Use list instead of arrays
// RETURN a list of kmer being t times in the dna
func list_kmer(dna string, k, t int) *list.List {
    res := list.New()
    my := make(map[string]int)
    var temp_wrd string

    for inc := 0; inc < len(dna); inc++ {
        if inc+k < len(dna) {
            temp_wrd = dna[inc : inc+k]
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
    for k, v := range my {
        if v >= t {
            res.PushBack(k)
        }
    }

    return res
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

// return the nb of occurrence of a pattern
func Kmer_occ(f_dna, f_pat []byte) int {
    var nb_occ = 0
    var i = 0

    for inc := 0; inc < len(f_dna); inc++ {
        for i+inc < len(f_dna) && f_dna[inc+i] == f_pat[i] {
            i++
            if i == len(f_pat) {
                nb_occ++
                i = 0
                break
            }
        }
        i = 0
    }
    return nb_occ
}
