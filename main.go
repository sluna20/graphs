package main

func main() {

}

type Node struct {
	Index         int
	AdjacentNodes map[int]*Node
}

// graph arcs are implicitly defined as those in which connected nodes belong to the graph
type Graph struct {
	Nodes map[int]*Node
}

func (g Graph) Inflow(node *Node) int {

	inflow := 0
	for _, v := range g.Nodes {
		for _, w := range v.AdjacentNodes {
			if w == node {
				inflow++
			}
		}
	}
	return inflow
}

func (g Graph) Outflow(node *Node) int {
	outflow := 0
	for _, v := range node.AdjacentNodes {
		for _, w := range g.Nodes {
			if v == w {
				outflow++
			}
		}
	}
	return outflow
}

func (g Graph) Netflow(node *Node) int {
	return g.Outflow(node) - g.Inflow(node)
}

func (g Graph) Degree(node *Node) int {
	return g.Outflow(node)
}

func (g Graph) IsComplete() bool {
	complete := true
	for _, v := range g.Nodes {
		if g.Degree(v) != len(g.Nodes)-1 {
			complete = false
			break
		}
	}
	return complete
}

func (g Graph) IsPath() bool {
	ends, transitions := 0, 0
	for _, v := range g.Nodes {
		if g.Degree(v) == 1 {
			ends++
		} else if g.Degree(v) == 1 {
			transitions++
		}
	}
	return ends == 2 && ends+transitions == len(g.Nodes)
}

func (g Graph) IsCycle() bool {
	transitions := 0
	for _, v := range g.Nodes {
		if g.Degree(v) == 0 {
			transitions++
		}
	}
	return transitions == len(g.Nodes)
}

func (g Graph) HasEuclidianPath() bool {
	ends := 0
	for _, v := range g.Nodes {
		if g.Degree(v) == 1 {
			ends++
		}
	}
	return ends == 0 || ends == 2
}
