/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"

*/
package main

import (
	"fmt"
)


func addBinary(a string, b string) string {
 var length int
	if (len(a) < len(b)){
		length = len(b) +1
	}else{
		length = len(a)+1
	}
	var temp = make([]byte,length)
	//var temp = make([]rune,length)
	c := 0
	i,j := len(a)-1,len(b)-1
	m := length -1
	for i>=0 || j>=0 ||c==1{
		if i >= 0{
			c +=  int(a[i]-'0')
			i--
		}
		if j >= 0{
			c += int(b[j]-'0')
			j--
		}
		temp[m] =byte(c%2 + '0')
		m--
		c /= 2

	}
	if temp[0] == 0{
		return string(temp[1:])
	}else{
		return string(temp[:])
	}
    
}
func main(){
 t := "11"
   td := "1"
	result := addBinary(t, td)
	fmt.Println(result)
}
