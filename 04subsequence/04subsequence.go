package main

import(
  "fmt"
  "math"
)
//n=10
//s=15
//a={5 1 3 5 10 7 4 9 2 8}
//-> 2 ( 5 + 10 )
func main() {
  var n, s, tmp,sum int
  var ary  []int
  fmt.Scanf("%d", &n)
  fmt.Scanf("%d", &s)

  for i:=0; i<n ;i++{
    fmt.Scanf("%d", &tmp)
    ary = append(ary, tmp)
  }

  var i_lower, i_higher, res int
  res = n+1
  i_lower = 0
  i_higher = 0
  sum = 0  
  for{
    for i_higher < n && sum < s {
      sum = sum +  ary[i_higher]
      i_higher++
    }
    if sum < s { 
      break
    }
    res = int(math.Min(float64(res), float64(i_higher - i_lower)))
    sum = sum - ary[i_lower]
    i_lower++
  }
  if(res > n){
    res = 0
  }
  fmt.Printf("%d",res)
}
