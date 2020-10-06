package main

import(
  "fmt" 
  "math" 
)

func main(){
  var n, k, tmp int
  var a []int
  var d []int
  
  fmt.Scanf("%d %d", &n, &k)
  for i:=0; i < n ; i++{
    fmt.Scanf("%d", &tmp)
    a = append(a, tmp)
    d = append(d, 1)
  }


  fmt.Println(AllCalc(n,k, a, d))  
  
}

func AllCalc(n,k int, a []int, d[] int)int{

  ans := 0
  for i := 1 ; i < n ; i++{
    d = calc(i, k, a, d)
    if d[i] > ans{
      ans = d[i]
    }
  }
  return ans
}

func max(a,b int)int{
  if a > b {
    return a
  }
  return b
}
func min(a,b int)int{
  if a < b {
    return a
  }
  return b
}

func calc(i, k int, a []int, d[] int)[]int{
  j := 1
  connected :=0
  plus_id := -1
  minus_id := -1
  //プラス側とマイナス側でいちばん近い接続先を確認
  for ;;{
    //fmt.Printf("%d %d %d %d %d %d %d¥n", i-j, i, a[i-j], a[i], k, minus_id, plus_id)
    
    if check_munus_connected(a[i - j], a[i], k) && minus_id < 0{
      connected++
      minus_id = i-j
    }else if check_plus_connected(a[i - j], a[i], k) && plus_id < 0{
      connected++
      plus_id = i-j
    }
    if connected == 2{
      break
    }
    j++
    if i-j < 0{
      break
    }
  }
  
  //fmt.Println(d)
  if plus_id < 0 && minus_id < 0{
    //fmt.Println("type a")
    d[i] = 1
  }else if plus_id < 0 {
    //fmt.Println("type b")
    d[i] = d[minus_id] + 1
  }else if minus_id < 0{
    //fmt.Println("type c")
    d[i] = d[plus_id] + 1
  }else if d[minus_id] < d[plus_id] {
    //fmt.Println("type d")
    d[i] = d[plus_id]+1
  }else if d[minus_id] > d[plus_id] {
    //fmt.Println("type d")
    d[i] = d[minus_id]+1
  }else{
    d[i] = d[minus_id]+1
  }
  
  //fmt.Println("aft:", d)
  return d
}

func check_munus_connected(a int, b int, k int)bool{
  
  //fmt.Println("m:", a, b, k)
  
  //1が5よりも −３以内か　−4
  if a - b <= 0 && math.Abs(float64(a - b)) <= float64(k) {
    //fmt.Println("t")
    return true
  }else{
    //fmt.Println("f")
    return false
  }
}

func check_plus_connected(a int, b int, k int)bool{
  //fmt.Println("p:", a, b, k)
  
  //1が5よりも +3 ~ 0以内か　-4
  if a - b >= 0 && math.Abs(float64(a - b)) <= float64(k) {
    //fmt.Println("t")
    return true
  }else{
    //fmt.Println("f")
    return false
  }
}
