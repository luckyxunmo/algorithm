/*
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Examples:
pattern = "abba", str = "dog cat cat dog" should return true.
pattern = "abba", str = "dog cat cat fish" should return false.
pattern = "aaaa", str = "dog cat cat dog" should return false.
pattern = "abba", str = "dog dog dog dog" should return false.
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters separated by a single space.

Credits:
Special thanks to @minglotus6 for adding this problem and creating all test cases.
*/

package main

import(
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
 	var patStrMap  = make(map[interface{}]int)
	words :=strings.Split(str," ")
	if len(pattern) != len(words){
		return false
	}
    for i := 0; i < len(words);i++{
    	if(tryPut(patStrMap,words[i],i) != tryPut(patStrMap,pattern[i],i)){
			return false
		}

	}
	return true
    
}
func tryPut(patStrMap map[interface{}]int, key interface{},value int) int  {
	if _,exist := patStrMap[key];exist{
		v := patStrMap[key]
		patStrMap[key] = value
		return v
	}else{
		patStrMap[key] = value
		return value
	}
}
func main()  {
	fmt.Println(wordPattern("abba","dog cat cat dog"))

}
