package main

import "fmt"

var dna = "atagagcatcagtcgatcgactagctacgtagtcagcatg"
var oriC = "atg"

// return the position of the p string in the s string, if p isn't in s then return -1
func check_str(s string, p string) int {
  var i = 0

  for inc := 0; inc < len(s); inc++ {
    for s[inc + i] == p[i] && inc + i < len(s){
      i++
      if i == len(p) {
        return inc + 1
      }
    }
    i = 0
  }
  return -1
}


func main() {
  fmt.Println(check_str(dna, oriC))
}
