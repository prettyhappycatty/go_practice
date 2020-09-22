package main

import(
  "fmt"
  "sort"
)
 
func main(){
  var n,tmp_x, tmp_y int
  var town_x []int
  var town_y []int
  var cnt int
  var group []int
  fmt.Scanf("%d", &n)
  group = make([]int, n)
  for i:= 0 ; i < n ; i++{
    fmt.Scanf("%d %d", &tmp_x, &tmp_y)
    town_x = append(town_x, tmp_x)
    town_y = append(town_y, tmp_y)
    group[i] = -1
  }
  
  
  //本編
  var paths [][]int
  var path []int
  var merged []int
  for l:=0 ; l< n; l++{
    if contains(merged, l){
      continue
    }
    path = canReach(l, path, town_x, town_y)
    if len(path) == 0{
      //fmt.Printf("if%d %d->", l)
      path = append(path, l)
    }
    paths = append(paths, path)
    merged = merge(merged, path)
    path = []int{}//path初期化
  }

  //fmt.Println("Paths:",paths)
  
  for i:=0; i<len(paths) ; i++{
    for j:=0; j<len(paths) ; j++{
      if i != j && SliceCommon(paths[i], paths[j]) { 
        //fmt.Println("Paths:",SliceCommon(paths[i], paths[j]), paths[i], paths[j])
        paths[i] = merge(paths[i], paths[j])
        paths[j] = []int{}
      }
    }
  }
  //fmt.Println("paths:",paths)
  
  for i:=0; i<len(paths) ; i++{
    for j:=0; j<len(paths[i]) ; j++{
      group[paths[i][j]] = len(paths[i])    
      //fmt.Println("group,cnt=", paths[i][j], len(paths[i]))
      cnt++
    }
    if cnt == n{
      break
    }
  }
  
  for l:=0 ; l< n; l++{
    fmt.Println(group[l])
  }
 
}
 
func canReach(start int, path []int, x []int, y []int)[]int{
  //fmt.Println("canReach1:",start, path)
  if contains(path, start) {
    //fmt.Println("contain")
    return path
  }
  path = append(path, start)
  //fmt.Println("canReach2:",start, path)
  for i := 0; i< len(x) ; i++ {
    if !contains(path, i) && can(x[start], y[start], x[i], y[i]) == true{
  //        fmt.Println("!contain & can=true")
          path = canReach(i, path, x, y)
    }
  }
  sort.Ints(path)
  return path
}
 
func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func can(x1, y1, x2, y2 int)bool{
  if((x1 > x2 && y1 > y2) ||
     (x1 < x2 && y1 < y2) ){
          return true
      }else{
        return false
      }
}

func merge(a []int, b []int)[]int{
  res := append(a, b...)  
  return SliceUnique(res)
}

func SliceUnique(target []int) (unique []int) {
    m := map[int]bool{}

    for _, v := range target {
        if !m[v] {
            m[v] = true
            unique = append(unique, v)
        }
    }

    return unique
}
        
func SliceCommon(target1, target2 []int)bool{
  merged := merge(target1, target2)
  if len(merged) < len(target1)+len(target2) {
    return true
  }
  return false
}
