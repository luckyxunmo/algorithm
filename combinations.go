////Combinations 组合问题
/*
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

For example,
If n = 4 and k = 2, a solution is:

[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

解法： 理解为 长度为N的数组 的组合 分为 包含第一个数字的组合加上不包含第一个数字的组合
*/
package main

import (
	"fmt"

)

func Combinations(source []string,num int)[]string  {
	if num == 1{
		result := make([]string,len(source))
		copy(result,source)
		return result
	}
	if len(source) <= num{
		var result string
		for _,k := range source{
			 result += k
			}
			return []string{result}
	}
	t1 := Combinations(source[1:],num-1)
	t := source[0]
     for i,_:= range t1{
     	t1[i] = t+t1[i]
	 }
	 fmt.Println("after t1 source is",source)
	t2 := Combinations(source[1:],num)
	return append(t1,t2...)

}


func main()  {
	temp :=[]string{"1","2","3","4","5"}
	fmt.Println(Combinations(temp,3))
}

