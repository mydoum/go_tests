package main

import (
    "flag"
    "fmt"
)

// flags
var reverse = flag.Bool("reverse", false, "reverse")
var r = flag.Bool("r", false, "reverse")

var frequent = flag.Bool("frequent", false, "frequent")
var f = flag.Bool("f", false, "frequent")

var help = flag.Bool("help", false, "help")
var h = flag.Bool("h", false, "help")

var occurrence = flag.Bool("occurrence", false, "occurrence")
var o = flag.Bool("o", false, "occurrence")

func fhelp() {
    fmt.Println("Usage: ./bio -rfho path_file")
}

func cmd_line() byte {
    flag.Usage = func() {
        fmt.Println("Usage: ./bio -rfho path_file")
    }

    flag.Parse()

    if *help || *h {
        fhelp()
    }

    if flag.NFlag() == 0 {
        fmt.Println("Flag missing")
    } else if flag.NFlag() > 1 {
        fmt.Println("Only one flag is required")
    }

    if len(flag.Args()) < 1 {
        fmt.Println("the path_file is missing")
        fhelp()
        return byte(1)
    } else if len(flag.Args()) > 1 {
        fmt.Println("Too many arguments!")
        fhelp()
        return byte(1)
    }

    if *reverse || *r {
        return 'r'
    } else if *frequent || *f {
        return 'f'
    } else if *occurrence || *o {
        return 'o'
    } else {
        fhelp()
        return byte(1)
    }
}
