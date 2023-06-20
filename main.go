package main

import (
	"fmt"
)

func main() {
	var testCompleteGraph Graph

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
	/*
		fmt.Println(testCompleteGraph)
		fmt.Println(testCompleteGraph.IsComplete())
		fmt.Println(testCompleteGraph.IsCycle())
		fmt.Println(testCompleteGraph.IsPath())
		fmt.Println(testCompleteGraph.HasEuclidianPath())
	*/
	n1 := testCompleteGraph.Nodes[0]
	n2 := testCompleteGraph.Nodes[1]
	num := testCompleteGraph.ShortestPath(n1, n2)
	fmt.Println(num)

	/*
		testPathGraph.addNode()
		testPathGraph.addNode()
		testPathGraph.addArc(testPathGraph.Nodes[0], testPathGraph.Nodes[1])
		testPathGraph.addArc(testPathGraph.Nodes[1], testPathGraph.Nodes[0])
		fmt.Println(testPathGraph)
		fmt.Println(testPathGraph.IsComplete())
		fmt.Println(testPathGraph.IsCycle())
		fmt.Println(testPathGraph.IsPath())
		fmt.Println(testPathGraph.HasEuclidianPath())
	*/
}

type Node struct {
	ID int
}

type Arc struct {
	From   Node
	To     Node
	Weight int
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
	nodeArcs := append(g.Arcs[from.ID], Arc{From: from, To: to, Weight: 2})
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

//DIJKSTRA

type Queue struct {
	distances []int
	prev      []Node
	nodes     []Node
	nodeIndex map[Node]int
	start     int
}

func (q *Queue) swap(i, j int) {
	q.distances[i], q.distances[j] = q.distances[j], q.distances[i]
	q.prev[i], q.prev[j] = q.prev[j], q.prev[i]
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
}

func (q *Queue) updateUp(i int) {

	current := i - q.start

	for {
		parent := ((current)+1)/2 - 1
		if current == 0 {
			break
		}
		if q.distances[q.start+current] < q.distances[q.start+parent] {
			//swap
			q.swap(q.start+current, q.start+parent)
			//move up
			current = parent
		} else {
			break
		}
	}

}

func (q *Queue) update(prev, node Node, dist int) {

	i, inQueue := q.nodeIndex[node]
	j := q.nodeIndex[prev]

	if !inQueue {
		i = len(q.nodes)

		q.distances = append(q.distances, dist)
		q.prev = append(q.prev, prev)
		q.nodes = append(q.nodes, node)
		q.nodeIndex[node] = i
		q.updateUp(i)
	} else if q.distances[j]+dist < q.distances[i] {

		q.distances[i] = dist
		q.prev[i] = prev
		q.nodes[i] = node

		q.updateUp(i)
	}
}

func (q *Queue) extract() {
	q.start++
}

func (g *Graph) ShortestPath(start Node, end Node) int {

	distances := make([]int, 1)
	prev := make([]Node, 1)
	nodes := make([]Node, 1)
	nodeIndex := make(map[Node]int, 1)

	nodeIndex[start] = 0
	q := Queue{
		distances: distances,
		prev:      prev,
		nodes:     nodes,
		nodeIndex: nodeIndex,
		start:     0,
	}

	i, inQueue := q.nodeIndex[end]
	for !inQueue || i > q.start {
		fmt.Println(q)
		//add or update neighbours of current node to queue
		arcs := g.Arcs[i]
		for _, a := range arcs {
			q.update(a.From, a.To, a.Weight)
		}
		q.extract()
		i, inQueue = q.nodeIndex[end]
	}
	return q.distances[i]
}
