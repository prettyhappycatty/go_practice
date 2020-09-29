package main

import(
  "fmt"
)
 
func main(){
  var n int
  n = 10000
  fmt.Println(n)
  for i:=0 ; i< n ; i++{
    fmt.Println(i+1, " ", n-i+1)
  }
}
