package algorithms_test

import (
	"bufio"
	"fmt"
	"graphs/algorithms"
	"os"
	"strconv"
	"testing"
)

func BenchmarkPrimSparse(b *testing.B) {

	file, err := os.Open("../ExampleGraphs/sparse.txt")
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
			seen := make([]bool, numOfVerticies)
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

			prims := algorithms.Prims{Verticies: numOfVerticies, Seen: seen, Graph: graph}

			prims.Construct(true)

			currentGraphIteration++
			numberOfGraphs--
		}

		// This is used in the case where not all of the test cases
		// in a file need to be read in
		if numberOfGraphs == 0 {
			break
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}

}

func BenchmarkPrimDense(b *testing.B) {

	file, err := os.Open("../ExampleGraphs/dense.txt")
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
			seen := make([]bool, numOfVerticies)
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

			prims := algorithms.Prims{Verticies: numOfVerticies, Seen: seen, Graph: graph}

			prims.Construct(true)

			currentGraphIteration++
			numberOfGraphs--
		}

		// This is used in the case where not all of the test cases
		// in a file need to be read in
		if numberOfGraphs == 0 {
			break
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}

}

// func BenchmarkPrim5nodes(b *testing.B) {

// 	file, err := os.Open("../ExampleGraphs/custom.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Close the file at the end
// 	defer file.Close()

// 	sc := bufio.NewScanner(file)
// 	sc.Split(bufio.ScanWords) // use scanwords

// 	for sc.Scan() {
// 		numberOfGraphs, _ := strconv.Atoi(sc.Text())
// 		currentGraphIteration := 1
// 		sc.Scan()

// 		for numberOfGraphs > 0 {
// 			numOfVerticies, _ := strconv.Atoi(sc.Text())
// 			graph := make([][]int, numOfVerticies)
// 			seen := make([]bool, numOfVerticies)
// 			index := 0
// 			sc.Scan()

// 			for i := 0; i < numOfVerticies; i++ {
// 				edgeConnectsTo := make([]int, numOfVerticies)
// 				for j := 0; j < numOfVerticies; j++ {
// 					cost, _ := strconv.Atoi(sc.Text())
// 					edgeConnectsTo[j] = cost
// 					sc.Scan()
// 				}

// 				graph[index] = edgeConnectsTo
// 				index++
// 			}

// 			prims := algorithms.Prims{Verticies: numOfVerticies, Seen: seen, Graph: graph}

// 			prims.Construct(true)

// 			currentGraphIteration++
// 			numberOfGraphs--
// 		}

// 		// This is used in the case where not all of the test cases
// 		// in a file need to be read in
// 		if numberOfGraphs == 0 {
// 			break
// 		}
// 	}

// 	if err := sc.Err(); err != nil {
// 		fmt.Println(err)
// 	}

// }

// func BenchmarkPrim8nodes(b *testing.B) {

// 	file, err := os.Open("../ExampleGraphs/custom2.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Close the file at the end
// 	defer file.Close()

// 	sc := bufio.NewScanner(file)
// 	sc.Split(bufio.ScanWords) // use scanwords

// 	for sc.Scan() {
// 		numberOfGraphs, _ := strconv.Atoi(sc.Text())
// 		currentGraphIteration := 1
// 		sc.Scan()

// 		for numberOfGraphs > 0 {
// 			numOfVerticies, _ := strconv.Atoi(sc.Text())
// 			graph := make([][]int, numOfVerticies)
// 			seen := make([]bool, numOfVerticies)
// 			index := 0
// 			sc.Scan()

// 			for i := 0; i < numOfVerticies; i++ {
// 				edgeConnectsTo := make([]int, numOfVerticies)
// 				for j := 0; j < numOfVerticies; j++ {
// 					cost, _ := strconv.Atoi(sc.Text())
// 					edgeConnectsTo[j] = cost
// 					sc.Scan()
// 				}

// 				graph[index] = edgeConnectsTo
// 				index++
// 			}

// 			prims := algorithms.Prims{Verticies: numOfVerticies, Seen: seen, Graph: graph}

// 			prims.Construct(true)

// 			currentGraphIteration++
// 			numberOfGraphs--
// 		}

// 		// This is used in the case where not all of the test cases
// 		// in a file need to be read in
// 		if numberOfGraphs == 0 {
// 			break
// 		}
// 	}

// 	if err := sc.Err(); err != nil {
// 		fmt.Println(err)
// 	}

// }


// func BenchmarkPrim10nodes(b *testing.B) {

// 	file, err := os.Open("../ExampleGraphs/custom3.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Close the file at the end
// 	defer file.Close()

// 	sc := bufio.NewScanner(file)
// 	sc.Split(bufio.ScanWords) // use scanwords

// 	for sc.Scan() {
// 		numberOfGraphs, _ := strconv.Atoi(sc.Text())
// 		currentGraphIteration := 1
// 		sc.Scan()

// 		for numberOfGraphs > 0 {
// 			numOfVerticies, _ := strconv.Atoi(sc.Text())
// 			graph := make([][]int, numOfVerticies)
// 			seen := make([]bool, numOfVerticies)
// 			index := 0
// 			sc.Scan()

// 			for i := 0; i < numOfVerticies; i++ {
// 				edgeConnectsTo := make([]int, numOfVerticies)
// 				for j := 0; j < numOfVerticies; j++ {
// 					cost, _ := strconv.Atoi(sc.Text())
// 					edgeConnectsTo[j] = cost
// 					sc.Scan()
// 				}

// 				graph[index] = edgeConnectsTo
// 				index++
// 			}

// 			prims := algorithms.Prims{Verticies: numOfVerticies, Seen: seen, Graph: graph}

// 			prims.Construct(true)

// 			currentGraphIteration++
// 			numberOfGraphs--
// 		}

// 		// This is used in the case where not all of the test cases
// 		// in a file need to be read in
// 		if numberOfGraphs == 0 {
// 			break
// 		}
// 	}

// 	if err := sc.Err(); err != nil {
// 		fmt.Println(err)
// 	}

// }

// func BenchmarkPrim18nodes(b *testing.B) {

// 	file, err := os.Open("../ExampleGraphs/18nodes.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Close the file at the end
// 	defer file.Close()

// 	sc := bufio.NewScanner(file)
// 	sc.Split(bufio.ScanWords) // use scanwords

// 	for sc.Scan() {
// 		numberOfGraphs, _ := strconv.Atoi(sc.Text())
// 		currentGraphIteration := 1
// 		sc.Scan()

// 		for numberOfGraphs > 0 {
// 			numOfVerticies, _ := strconv.Atoi(sc.Text())
// 			graph := make([][]int, numOfVerticies)
// 			seen := make([]bool, numOfVerticies)
// 			index := 0
// 			sc.Scan()

// 			for i := 0; i < numOfVerticies; i++ {
// 				edgeConnectsTo := make([]int, numOfVerticies)
// 				for j := 0; j < numOfVerticies; j++ {
// 					cost, _ := strconv.Atoi(sc.Text())
// 					edgeConnectsTo[j] = cost
// 					sc.Scan()
// 				}

// 				graph[index] = edgeConnectsTo
// 				index++
// 			}

// 			prims := algorithms.Prims{Verticies: numOfVerticies, Seen: seen, Graph: graph}

// 			prims.Construct(true)

// 			currentGraphIteration++
// 			numberOfGraphs--
// 		}

// 		// This is used in the case where not all of the test cases
// 		// in a file need to be read in
// 		if numberOfGraphs == 0 {
// 			break
// 		}
// 	}

// 	if err := sc.Err(); err != nil {
// 		fmt.Println(err)
// 	}

// }
