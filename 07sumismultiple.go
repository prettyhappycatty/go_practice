package main

import(
  "fmt"
//  "math"
  "sort"
)

func main(){
  var n int
  fmt.Scanf("%d", &n)
  
  if(n==1){
      fmt.Println(1)
      return
  }
  
  //約数
  var pf []int
  var primemap map[int]int
  pf = prime_factor(n)
  fmt.Println("prime_facor", pf)
  fmt.Printf("%d",len(pf))
  primemap = prime_map(pf)
  
  fmt.Println(primemap)
  
}

//約数を求める
func devisor(n int)[]int{
  var res []int
  for i:=2 ; i*i < n ; i++{
    if ( n % i == 0){
      res = append(res, i)
      res = append(res, n/i)
    }
  }
  if n != 1 {
    res = append(res, 1)
    res = append(res, n)
  }
  sort.Ints(res)
  return res
}

//素因数分解
func prime_factor(n int)[]int{
  var res []int
  var rest int
  rest = n
  Label:
  for i:=2 ; i*i < n ; i++{
    for ;;{
      if rest % i == 0 {
        res = append(res, i)
        rest = rest / i
      }else{
        continue Label
      }
    } 
  }
  sort.Ints(res)
  return res
}

func prime_map(prime_factor []int)map[int]int{
  m := map[int]int{}
  for i:=0 ; i < len(prime_factor) ; i ++{
    if _, ok := m[prime_factor[i]] ; ok{
      m[prime_factor[i]]++
    }else{
      m[prime_factor[i]] = 1
    }
  }
    return m
}
