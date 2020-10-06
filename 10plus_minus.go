//X+Y = A
//X-Y = B
//A+B = 2X  X=(A+B)/2
//A-B = 2Y  Y=(A-B)/2

package main

import(
  "fmt"
)

func main(){
  var a,b int
  fmt.Scanf("%d %d", &a, &b)
  
  fmt.Printf("%d %d", (a+b)/2, (a-b)/2)
  
  
}

