package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	v int // v 个节点
	adj map[int][]int

}

func (g *Graph) init(v int)  {
	g.v = v
	g.adj = make(map[int][]int)

}
func (g *Graph) addEdage(s,t int)  {
	if g.adj[s] == nil{
		g.adj[s] = make([]int,0)
	}
	if g.adj[t] == nil{
		g.adj[t] = make([]int,0)
	}
	g.adj[s] =  append(g.adj[s],t)
	g.adj[t] =  append(g.adj[t],s)

}

func (g *Graph)Bfs(s int,t int)  {
	visited := make(map[int]bool)
	pre := make(map[int]int)
	visited[s] = true
	queue := list.New()
	queue.PushBack(s)

	for{
		if visited[t] == true{
			break
		}
		va := queue.Front().Value
		node :=va .(int)
		fmt.Println("node is",node)
		pengyou := g.adj[node]
	     for i := 0; i < len(pengyou); i++{
	     	peng := pengyou[i]
	     	if !visited[peng]{
	     		queue.PushBack(peng)
	     		fmt.Printf("%d in queue\n",peng)
	     		visited[peng] = true
	     		pre[peng] = node
			}

		 }
		 queue.Remove(queue.Front())
	}
	zhengxuPrint(pre,s,t)

}

func (g *Graph)DFSUseStatck(s,t int)  {
	visited := make(map[int]bool)
	stack:= list.New()
	prev := make(map[int]int)
    stack.PushBack(s)
	visited[s] = true
	for{
		node := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		fmt.Println("node is:",node)
		for _,v := range g.adj[node]{
			if v == t{
				prev[v] = node
				goto end
			}
			fmt.Printf("node:%v, v:%v\n",node,v)
			if !visited[v]{
				visited[v] =  true
				prev[v] = node
				stack.PushBack(v)
			}

		}
	}
	end:
		zhengxuPrint(prev,s,t)

}

// 从前往后打印
func zhengxuPrint(pre map[int]int,s,t int)  {
	_,ok := pre[t]
	if ok && t != s{
		fmt.Printf("now t:%v,s:%v\n",t,s)
		zhengxuPrint(pre,s,pre[t])
	}
	fmt.Println(t)


}

// 从后往前打印
func nixuPrint(pre map[int]int,s,t int)  {
	fmt.Println(t)
	m := pre[t]
	for m != s{
		fmt.Println(m)
		n := pre[m]
		m = n
	}
	fmt.Println(m)
}

func main() {
	p := Graph{}
	p.init(10)
	p.addEdage(1,3)
	p.addEdage(1,2)
	p.addEdage(3,4)
	p.addEdage(3,5)
	p.addEdage(2,5)
	p.addEdage(2,6)
	p.addEdage(5,7)
	p.addEdage(7,8)
	p.addEdage(7,9)
	p.addEdage(6,10)
	//p.Bfs(1,7)
	p.DFSUseStatck(1,7)
}
