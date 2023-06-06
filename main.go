package main

import "fmt"

func main() {
	var testCompleteGraph, testPathGraph Graph

	for i := 0; i < 5; i++ {
		testCompleteGraph.addNode()
	}
	for i, n := range testCompleteGraph.Nodes {
		for j, m := range testCompleteGraph.Nodes {
			if i != j {
				testCompleteGraph.addArc(n, m)
			}
		}
	}
	fmt.Println(testCompleteGraph)
	fmt.Println(testCompleteGraph.IsComplete())
	fmt.Println(testCompleteGraph.IsCycle())
	fmt.Println(testCompleteGraph.IsPath())
	fmt.Println(testCompleteGraph.HasEuclidianPath())

	testPathGraph.addNode()
	testPathGraph.addNode()
	testPathGraph.addArc(testPathGraph.Nodes[0], testPathGraph.Nodes[1])
	testPathGraph.addArc(testPathGraph.Nodes[1], testPathGraph.Nodes[0])
	fmt.Println(testPathGraph)
	fmt.Println(testPathGraph.IsComplete())
	fmt.Println(testPathGraph.IsCycle())
	fmt.Println(testPathGraph.IsPath())
	fmt.Println(testPathGraph.HasEuclidianPath())
}

type Node struct {
	ID int
}

type Arc struct {
	From int
	To   int
}

type Graph struct {
	Nodes []Node
	Arcs  [][]Arc
}

func (g *Graph) addNode() {
	g.Nodes = append(g.Nodes, Node{ID: len(g.Nodes)})
	g.Arcs = append(g.Arcs, make([]Arc, 0))
}

func (g *Graph) addArc(from Node, to Node) {
	nodeArcs := append(g.Arcs[from.ID], Arc{From: from.ID, To: to.ID})
	g.Arcs[from.ID] = nodeArcs
}

func (g Graph) Degree(n Node) int {
	return len(g.Arcs[n.ID])
}

func (g Graph) IsComplete() bool {
	complete := true
	for _, n := range g.Nodes {
		if g.Degree(n) != len(g.Nodes)-1 {
			complete = false
			break
		}
	}
	return complete
}

func (g Graph) IsPath() bool {
	ends, transitions := 0, 0
	for _, n := range g.Nodes {
		if g.Degree(n) == 1 {
			ends++
		} else if g.Degree(n) == 2 {
			transitions++
		}
	}
	return ends == 2 && ends+transitions == len(g.Nodes)
}

func (g Graph) IsCycle() bool {
	transitions := 0
	for _, n := range g.Nodes {
		if g.Degree(n) == 0 {
			transitions++
		}
	}
	return transitions == len(g.Nodes)
}

func (g Graph) HasEuclidianPath() bool {
	odd := 0
	for _, n := range g.Nodes {
		if g.Degree(n)%2 == 1 {
			odd++
		}
	}
	return odd == 0 || odd == 2
}
