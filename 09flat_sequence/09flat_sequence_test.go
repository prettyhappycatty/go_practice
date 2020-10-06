package main

import (
    "testing"
)
func TestAllCalc(t *testing.T){
  var a, d []int
  var k, n int
  a = []int{1,5,4,3,8,6,9,7,2,4}
  k = 3
  n = 10
  d = make([]int,n)
  for i := 0 ; i<n ; i++{
    d = append(d, 1)
  }
  result := AllCalc(n, k, a, d)
  expect := 7

  if result != expect {
    t.Error("\n実際： ", result, "\n理想： ", expect)
  }

}
