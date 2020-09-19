package main

import(
  "fmt"
  "sort"
)

func main() {
  var n, tmp int
  var ary []int
  fmt.Scanf("%d", &n)
  
  for i:=0; i<n ;i++{
    fmt.Scanf("%d", &tmp)
    ary = append(ary, tmp)
  }
  sort.Slice(ary,func(i, j int) bool  {
    return ary[i] > ary[j]
  })

  var i,j,k int 
  for i= 0 ; i < n ; i++ {
    for j= i+1 ; j < n ; j++ {
      for k = j+1 ; k < n ; k++ {
        if(ary[i] < ary[j] + ary[k]){
          fmt.Println(ary[i] + ary[j] + ary[k])
          return
        }
      }
    }
  }
  fmt.Println("0")
}
