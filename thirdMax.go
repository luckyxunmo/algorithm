/*
Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
*/
package thirdMax

import (
	"container/heap"
	"math"
)

func thirdMax(nums []int) int {
    	if len(nums) ==1 {
		return nums[0]
	}
	if len(nums) == 2{
		return 	int(math.Max(float64(nums[0]),float64(nums[1])))
	}
	a :=IntHeap{}
	a = append(a,nums[0:3]...)
	h := &a
	heap.Init(h)
	for j := 3; j < len(nums);j++{
		if (*h)[0] < nums[j]{
			(*h)[0],nums[j] = nums[j],(*h)[0]
			heap.Fix(h,0)
		}
	}

	return (*h)[0]     
                    
}
type IntHeap []int

func (h IntHeap) Len() int{
	return len(h)
}
func (h IntHeap)Less(i,j int) bool  {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i,j int)  {
	h[i],h[j] = h[j],h[i]
}

func (h *IntHeap)Push(x interface{})  {
	*h = append(*h,x.(int))
}
func (h *IntHeap) Pop() interface{}  {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}
