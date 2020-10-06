package main

import(
  "fmt"
)

type UnionFind struct{
  n int
  par []int
  rank []int
}

func (u UnionFind) init() {
  u.par = []int{}
  u.rank = []int{}
  for i:= 0 ; i < u.n ; i++{
    par = append(u.par, i)
    rank = append(u.rank, 0)
  }
}

func (u UnionFind) find(x int) { 
  if u.par[x] == x{
    return x
  }else{
    u.par[x] = find(u.par[x])
    return u.par[x]
  }
}

func (u UnionFind) unite(x int, y int) {
  x = find(x)
  y = find(y)
  if x == y{
    return
  }
  if u.rank[x] < u.rank[y]{
    par[x] =y
  }else{
    if u.par[x] == x{
    par[y] =x
    if u.rank[x] == u.rank[y]{
      rank[x]++
    }
}

func (u UnionFind) is_same(x int, y int) {
  return find(x) == find(y)
}

func main(){
  var n,tmp_x, tmp_y int
  var town_x []int
  var town_y []int
  var cnt int
  var group []int
  
  fmt.Println("1:",time.Now())
  fmt.Scanf("%d", &n)
  group = make([]int, n)
  for i:= 0 ; i < n ; i++{
    fmt.Scanf("%d %d", &tmp_x, &tmp_y)
    town_x = append(town_x, tmp_x)
    town_y = append(town_y, tmp_y)
    group[i] = -1
  }
  //Union-Find-Tree
  var per []int
  var rank []int

  for i:= 0 ; i < n ; i++{
    per = append(per, i)
    rank = append(rank, 0)
  }

}

func init(par, rank int)
