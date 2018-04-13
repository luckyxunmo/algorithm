/*
 Permutation Sequence

The set [1,2,3,…,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order,
We get the following sequence (ie, for n = 3):

"123"
"132"
"213"
"231"
"312"
"321"

解法：采用了交换元素+DFS的方法来求解。理解为分别将长度为n的数组的每一个数与数组的第一个数进行交换，然后对后面的n-1子数组进行全排列
*/
package main

import "fmt"

var res = make([][]string,0)
func Swap(source []string,i,j int)  {
	source[i],source[j] = source[j],source[i]
}

func Permutation(source []string,i,len int)  {
	if i == len-1{
		var list []string
		for j := 0;j<len;j++{
			list = append(list,source[j])
		}
		res = append(res,list)
		return
	}

   for m :=i;m<len;m++{
   	 Swap(source,i,m)
   	 Permutation(source,i+1,len)
   	 Swap(source,i,m)
   }

}

func main()  {
	temp :=[]string{"1","2","3"}
	Permutation(temp,0,len(temp))
	fmt.Println(res)
}
