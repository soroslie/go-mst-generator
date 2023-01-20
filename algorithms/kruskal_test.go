package algorithms_test

import (
	"bufio"
	"fmt"
	"graphs/algorithms"
	edge "graphs/edge"
	"os"
	"strconv"
	"testing"
)

func BenchmarkKruskal(b *testing.B) {

	file, err := os.Open("../ExampleGraphs/custom.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Close the file at the end
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords) // use scanwords

	for sc.Scan() {
		numberOfGraphs, _ := strconv.Atoi(sc.Text())
		currentGraphIteration := 1
		sc.Scan()

		for numberOfGraphs > 0 {
			numOfVerticies, _ := strconv.Atoi(sc.Text())
			graph := make([][]int, numOfVerticies)
			index := 0
			sc.Scan()

			for i := 0; i < numOfVerticies; i++ {
				edgeConnectsTo := make([]int, numOfVerticies)
				for j := 0; j < numOfVerticies; j++ {
					cost, _ := strconv.Atoi(sc.Text())
					edgeConnectsTo[j] = cost
					sc.Scan()
				}

				graph[index] = edgeConnectsTo
				index++
			}

			edges := edge.MakeEdges(graph, numOfVerticies)

			kruskal := algorithms.Kruskals{Verticies: numOfVerticies, Edges: edges}

			kruskal.Construct(true)

			currentGraphIteration++
			numberOfGraphs--
		}

		if numberOfGraphs == 0 {
			break
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}

}
