package main

import(
  "fmt"
)

func main(){
  var d, t, s int
  fmt.Scanf("%d %d %d", &d, &t, &s)
  
  if(check(d, s, t)){
    fmt.Printf("Yes")
  }else{
    fmt.Printf("No")
  }
  
}


func check(d, s, t int)bool{
  
  if(d <= s * t){
    return true
  }else{
    return false
  }

}
