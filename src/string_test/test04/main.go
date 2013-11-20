package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "string_test/test04/binfo"
    "strings"
)

// variables and structures
var file_name []byte

// flags
var reverse = flag.Bool("reverse", false, "reverse")
var r = flag.Bool("r", false, "reverse")

var frequent = flag.Bool("frequent", false, "frequent")
var f = flag.Bool("f", false, "frequent")

var help = flag.Bool("help", false, "help")
var h = flag.Bool("h", false, "help")

var occurrence = flag.Bool("occurrence", false, "occurrence")
var o = flag.Bool("o", false, "occurrence")

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
    }
}

func main() {
    flag.Usage = func() {
        fmt.Println("usage: bio < -rf > path_file")
    }

    flag.Parse()
    args := flag.Args()

    if *help || *h {
        //TODO
    }

    if flag.NFlag() == 0 {
        fmt.Println("Flag missing")
    } else if flag.NFlag() > 1 {
        fmt.Println("Only one flag is required")
    }

    if len(flag.Args()) < 1 {
        fmt.Println("the path_file is missing")
        return
    } else if len(flag.Args()) > 1 {
        fmt.Println("Too many arguments!")
        return
    } else {
        file_name = []byte(args[0])
    }

    if *reverse || *r {
        apply(file_name, 'r')
    } else if *frequent || *f {
        apply(file_name, 'f')
    } else if *occurrence || *o {
        apply(file_name, 'o')
    }

    return
}
