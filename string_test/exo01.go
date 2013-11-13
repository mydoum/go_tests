package main

import "fmt"

var dna = "atagagcatcagtcgatcgactagctacgtagtcagcatg"
var oriC = "atg"

// return the position of the p string in the s string, if p isn't in s then return 0
func check_str(s string, p string) int {
  var inc int
  var i = 0

  for inc := 0; inc < len(s); inc++ {
    for s[inc + i] == p[inc + i] {
      i++
      if i == len(p) {
        return inc
      }
    }
    i = 0
  }
  return 0
}

func main() {
  fmt.Println(check_str(dna, oriC))
}
