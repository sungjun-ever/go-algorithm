package main

import "fmt"

type Edge struct {
	To     int
	Weight int
}

// ListGraph는 인접 리스트 방식의 그래프
type ListGraph struct {
	Vertices int
	AdjList  map[int][]Edge
}

func NewListGraph(v int) *ListGraph {
	return &ListGraph{
		Vertices: v,
		AdjList:  make(map[int][]Edge),
	}
}

// 방향성 및 가중치 간성 추가
func (g *ListGraph) AddEdge(from, to, weight int) {
	g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})
}

func (g *ListGraph) Print() {
	fmt.Println("--- 인접 리스트 그래프 ---")
	for i := 0; i < g.Vertices; i++ {
		fmt.Printf("정점 %d: ", i)
		for _, edge := range g.AdjList[i] {
			fmt.Printf("-> %d(가중치: %d) ", edge.To, edge.Weight)
		}
		fmt.Println()
	}
}

type MatrixGraph struct {
	Vertices int
	Matrix   [][]int
}

func NewMatrixGraph(v int) *MatrixGraph {
	matrix := make([][]int, v)
	for i := range matrix {
		matrix[i] = make([]int, v)
	}
	return &MatrixGraph{
		Vertices: v,
		Matrix:   matrix,
	}
}

func (g *MatrixGraph) AddEdge(from, to, weight int) {
	g.Matrix[from][to] = weight
}

func (g *MatrixGraph) Print() {
	fmt.Println("--- 인접 행렬 그래프 ---")
	for i := 0; i < g.Vertices; i++ {
		for j := 0; j < g.Vertices; j++ {
			fmt.Printf("%2d ", g.Matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {

}
