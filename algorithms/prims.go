package algorithms

import "fmt"

var INF int = 9999999

type Prims struct {
	Edges     int
	Verticies int
	Graph     [][]int
	Seen      []bool
	Total     int
}

func (p *Prims) Construct(isTest bool) {
	p.Seen[0] = true
	p.Edges = 0
	p.Total = 0

	for p.Edges < p.Verticies-1 {
		min := INF
		src, dest := 0, 0
		for i := 0; i < p.Verticies; i++ {
			if p.Seen[i] {
				for j := 0; j < p.Verticies; j++ {
					if !p.Seen[j] && p.Graph[i][j] != 0 {
						if min > p.Graph[i][j] {
							src, dest, min = i, j, p.Graph[i][j]
						}
					}
				}
			}
		}

		if !isTest {
			fmt.Printf("%d - %d | %d\n", src, dest, p.Graph[src][dest])
		}
		p.Seen[dest] = true
		p.Total += min
		p.Edges++
	}
}
