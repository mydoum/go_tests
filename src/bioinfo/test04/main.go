package main

import (
    "bioinfo/test04/binfo"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
    "time"
)

// variables and structures
var file_name []byte

// Put dna in a buffer from the path_file and apply the correct fct
func apply(Text_name []byte, mod byte) {
    buffer, err := ioutil.ReadFile(string(Text_name))

    if err != nil {
        log.Fatal(err)
    }

    var list []string = strings.Split(string(buffer), "\n")

    // allow to escape the first array if we have a Format1 type
    i := 0
    if list[0] == "Input" {
        i = 1
    }

    dna := list[i]
    if mod == 'f' {
        k, _ := strconv.Atoi(list[i+1])
        binfo.Most_kmer(dna, k)
    } else if mod == 'r' {
        binfo.Compl_strand(binfo.Reverse(dna))
    } else if mod == 'o' {
        binfo.Pat_positions(list[i], list[i+1])
    } else if mod == 'c' {
        params := strings.Split(list[i+1], " ")
        k, _ := strconv.Atoi(params[0])
        L, _ := strconv.Atoi(params[1])
        t, _ := strconv.Atoi(params[2])
        binfo.Clump(list[i], k, L, t)
    }
}

func main() {
    flg := cmd_line()
    args := flag.Args()

    if flg == byte(1) {
        return
    } else {
        file_name = []byte(args[0])
    }

    t0 := time.Now()

    apply(file_name, flg)

    t1 := time.Now()
    fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

    return
}
