package main

import "fmt"

type Item interface {
}

// 组成图的顶点
type Node struct {
	Value  Item
	IsOpen bool
}

func (n *Node) Open() {
	n.IsOpen = true
}

// 定义一个图的结构, 图有顶点与边组成 V  E
type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// 添加节点
func (g *ItemGraph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// 添加边
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	// 无向图
	g.Edges[*n1] = append(g.Edges[*n1], n2) // 设定从节点n1到n2的边
	g.Edges[*n2] = append(g.Edges[*n2], n1) // 设定从节点n2到n1的边
}

func (g *ItemGraph) String() {
	s := ""

	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + "->"
		near := g.Edges[*g.Nodes[i]]

		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}

	fmt.Println(s)
}

func (g *ItemGraph) Go(n1, n2 *Node) bool {
	for _, near := range g.Edges[*n1] {
		if near.String() == n2.String() && n2.IsOpen {
			return true
		}
	}
	return false
}

var g ItemGraph

func fillGraph() {
	nA := Node{Value: "A"}
	nB := Node{Value: "B"}
	nC := Node{Value: "C"}
	nD := Node{Value: "D"}
	nE := Node{Value: "E"}
	nF := Node{Value: "F"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}

func main() {
	fillGraph()
	g.String()

	nA := &Node{Value: "B"}
	nB := &Node{Value: "A"}

	fmt.Println(g.Go(nA, nB))
}
