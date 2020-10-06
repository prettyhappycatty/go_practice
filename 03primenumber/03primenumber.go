package main

import(
  "fmt"
  "math"
)

func main(){
  var n int
  var ans []int
  fmt.Scanf("%d", &n)
  ans = prime(n)
  fmt.Println(ans)
}

func prime(num int)[]int{
  var a []int
  a = append(a, 2)  
  for i:=3;i<=num;i+=2 {
        k:=0
        for j:=3;j<=int(math.Sqrt(float64(i)));j+=2{   
            if(i%j==0){
                k=1
                break
           }
        }
        if(k==0){
          a = append(a, i)
        }
    }

  return a
}
