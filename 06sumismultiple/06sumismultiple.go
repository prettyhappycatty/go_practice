package main

import(
  "fmt"
  "math"
)

func main(){
  var n int64
  fmt.Scanf("%d", &n)
  
  if(n==1){
      fmt.Println(1)
      return
  }
  
  var start int64
  start = int64(math.Sqrt(float64(n)))
  //fmt.Printf("%d", start)
  for i:=start+1 ; i<n*n ; i++{
    if( i * (i+1) /2  % n == 0){
      fmt.Printf("%d %d", i, i * (i+1) /2)
      return
  	}
  }
}

