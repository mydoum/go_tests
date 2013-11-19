package main

import (
    "string_test/test03/kmer"
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "strings"
    "strconv"
)

// variables and structures
var k []byte
var file_name []byte

func frq_wrd_with_k(Text_name []byte, k int) []string {
    return []string{"a", "b"}
}

// Take a text_name and return a list of most_kmer
func frq_wrd(Text_name []byte) []string {
    buffer, err := ioutil.ReadFile(string(Text_name))

    if err != nil {
            log.Fatal(err)
    }

    var list []string = strings.Split(string(buffer), "\n")

    // allow to escape the first array if we have a Format1 type
    i := 0
    if list[0]=="Input" {
        i = 1
    }

    dna := list[i]
    k, err := strconv.Atoi(list[i+1])

    // init mapping
    kmer.Most_kmer(dna, k)
    //my := make(map[string]int)

    return []string{"o", "i"}
}

func main() {
    if len(os.Args) < 2 || len(os.Args) > 3 {
        fmt.Println("frw_wrd_pb filename [k int]")
        return
    }

    if len(os.Args) == 3 {
        file_name = []byte(os.Args[1])
        k = []byte(os.Args[2])
        // call frq_wrd_with_k
        // Print the list of words
        return
    }    else {
        file_name = []byte(os.Args[1])
        frq_wrd(file_name)
        return
    }
}
