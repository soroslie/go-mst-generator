package algorithms

import (
	"fmt"
	dsu "graphs/dsu"
	edge "graphs/edge"
	"sort"
)
type Kruskals struct {
	Verticies int
	Edges     []edge.Edge
	Total     int
}

func (k *Kruskals) Construct(isTest bool) {
	index := 0

	sort.Sort(edge.Edges(k.Edges))

	subsets := make([]dsu.Subset, k.Verticies)
	for i := 0; i < k.Verticies; i++ {
		subsets[i] = dsu.Subset{Parent: i, Children: 0}
	}

	result := make([]edge.Edge, k.Verticies-1)
	index = 0
	i := 0

	for index < k.Verticies-1 {
		if i >= len(k.Edges) {
			break
		}

		next := k.Edges[i]
		x := dsu.Find(subsets, next.Src)
		y := dsu.Find(subsets, next.Dest)

		if x != y {
			result[index] = next
			dsu.Union(subsets, x, y)
			index++
		}

		i++
	}

	for _, anEdge := range result {
		if !isTest {
			fmt.Printf("%d - %d | %d\n", anEdge.Src, anEdge.Dest, anEdge.Cost)

		}
		k.Total += anEdge.Cost
	}
}

