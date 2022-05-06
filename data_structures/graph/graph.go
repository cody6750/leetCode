package main

import (
	"log"
)

type Graph struct {
	Vertecies []*Vertex
	// Vertex Value and Index
	Duplicates map[int]*Vertex
}

type Vertex struct {
	Value int
	Edges []*Vertex
	// Edge Value and Index
	Duplicates map[int]*Vertex
}

func (g *Graph) Insert(v int) {

	vertex := &Vertex{Value: v, Duplicates: map[int]*Vertex{}}
	if _, exist := g.Duplicates[vertex.Value]; exist {
		log.Printf("Duplicate found %v", vertex.Value)
		return
	}
	g.Vertecies = append(g.Vertecies, vertex)
	g.Duplicates[vertex.Value] = vertex
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.Contains(from)
	toVertex := g.Contains(to)
	if fromVertex != nil || toVertex != nil {
		log.Printf("From %v does not exist in the list of verticies", from)
		return
	}

	if _, exist := fromVertex.Duplicates[toVertex.Value]; exist {
		return
	}
	fromVertex.Edges = append(fromVertex.Edges, toVertex)
	fromVertex.Duplicates[toVertex.Value] = toVertex

	return
}

func (g *Graph) Contains(v int) *Vertex {
	if len(g.Vertecies) == 0 {
		log.Print("Graph verticies is empty")
		return nil
	}
	for _, vertex := range g.Vertecies {
		if vertex.Value == v {
			return vertex
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, g := range g.Vertecies {
		log.Print(g.Value, g.Edges)
	}
}

func (g *Graph) CreateAllEdges() {
	for i := 0; i < len(g.Vertecies); i++ {
		for j := 0; j < len(g.Vertecies); j++ {
			g.AddEdge(i, j)
		}
	}
}

func (g *Graph) BFS(v int) bool {
	if len(g.Vertecies) == 0 {
		return false
	}
	queue := []*Vertex{}
	queue = append(queue, g.Vertecies[0])
	visited := make(map[*Vertex]struct{})
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = struct{}{}
		if current.Value == v {
			return true
		}
		for _, edge := range current.Edges {
			if _, exist := visited[edge]; exist {
				continue
			}
			queue = append(queue, edge)
		}
	}
	return false
}

func (g *Graph) DFS(v int) bool {
	return g.DFSRecursive(v, g.Vertecies[0], make(map[*Vertex]int))
}

func (g *Graph) DFSRecursive(v int, vertex *Vertex, visited map[*Vertex]int) bool {
	var result bool
	if len(g.Vertecies) == 0 {
		return result
	}
	if vertex.Value == v {
		return true
	}
	for _, edge := range g.Vertecies {
		if _, exist := visited[edge]; exist {
			continue
		}
		visited[edge] = edge.Value
		result = g.DFSRecursive(v, edge, visited)
	}
	return result
}
func initGraph() {
	g := Graph{}
	g.Duplicates = make(map[int]*Vertex)
	g.Insert(0)
	g.Insert(0)
	g.Insert(1)
	g.Insert(2)
	g.Insert(3)
	g.Insert(4)
	g.Insert(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 5)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 1)
	g.Print()
	log.Print(g.DFS(0))
}
