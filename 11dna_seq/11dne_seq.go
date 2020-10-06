package main

import(
  "fmt"
)

func main(){
  var n int
  var s string
  var cnt_a, cnt_c, cnt_g, cnt_t int
  fmt.Scanf("%d %s", &n, &s)
  
  ans := 0
  for i:=0; i<n ; i++{
    cnt_a = 0
    cnt_t = 0
    cnt_c = 0
    cnt_g = 0
    for j:=i; j<n; j++{
      if s[j] == 'A' {
        cnt_a++
      }else if s[j] == 'C'{
        cnt_c++
      }else if s[j] == 'G'{
        cnt_g++
      }else if s[j] == 'T'{
        cnt_t++
      }
      if cnt_a == cnt_t && cnt_c == cnt_g{
        ans++ 
      }
    }
  }
  fmt.Println(ans)
}
