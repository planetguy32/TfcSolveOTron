package main

import (
	"fmt"

	"github.com/fzipp/astar"
)

type Button int

func reverse(in map[int]Button) map[Button]int {
	result := map[Button]int{}
	for k, v := range in {
		result[v] = k
	}
	return result
}

const (
	// The start and end states have t
	startLBP   Button = -2
	endLBP     Button = -1
	invalidLBP Button = -3
)

var deltaToPosition = map[int]Button{
	0:   0,
	-3:  1,
	-6:  2,
	2:   3,
	7:   4,
	-9:  5,
	-15: 6,
	13:  7,
	16:  8,
}

var positionToDelta = reverse(deltaToPosition)

type Graph struct{}

type Node struct {
	lastButtonPress Button
	position        int
}

func (g Graph) Neighbours(n Node) []Node {
	if n.lastButtonPress == endLBP {
		return []Node{}
	}
	results := []Node{
		{
			lastButtonPress: endLBP,
			position:        n.position,
		},
	}
	for i, delta := range positionToDelta {
		newValue := Node{
			lastButtonPress: i,
			position:        n.position + delta,
		}
		if newValue.position >= 0 && newValue.position <= 150 {
			results = append(results, newValue)
		}
	}
	return results
}

type ButtonPattern struct {
	b Button
	x int
}

func rle(in []Node) []ButtonPattern {
	var result []ButtonPattern
	if len(in) == 0 {
		return nil
	}
	buildingPattern := ButtonPattern{
		b: in[0].lastButtonPress,
		x: 1,
	}
	for i := 1; i < len(in); i++ {
		if in[i].lastButtonPress != buildingPattern.b {
			result = append(result, buildingPattern)
			buildingPattern = ButtonPattern{
				b: in[i].lastButtonPress,
				x: 0,
			}
		}
		buildingPattern.x++
	}
	return result
}

func main() {
	realCost := func(i, j Node) float64 {
		if j.lastButtonPress == -1 {
			return 0
		}
		// Slightly prefer repeating the same button
		if j.lastButtonPress != i.lastButtonPress {
			return 1.0001
		}
		return 1
	}
	// Can't think of a working heuristic off the top of my head.
	heuristicCost := func(i, j Node) float64 {
		return 0
	}
	var graph astar.Graph[Node]
	graph = Graph{}

	fmt.Println("Enter current position: ")
	var curPos int
	fmt.Scanf("%d", &curPos)

	fmt.Println("Enter final position: ")
	var finalPos int
	fmt.Scanf("%d", &finalPos)

	fmt.Println("\nEnter last button positions: ")
	var l1 int
	var l2 int
	var l3 int
	fmt.Scanf("%d %d %d", &l1, &l2, &l3)
	goalIdx := finalPos - positionToDelta[Button(l1)] - positionToDelta[Button(l2)] - positionToDelta[Button(l3)]

	startNode := Node{
		position:        curPos,
		lastButtonPress: startLBP,
	}

	goalNode := Node{
		position:        goalIdx,
		lastButtonPress: endLBP,
	}

	var path []Node
	path = astar.FindPath(graph, startNode, goalNode, realCost, heuristicCost)

	rlePath := rle(path)

	fmt.Printf("\n\n")
	for _, a := range rlePath {
		if a.b != startLBP && a.b != endLBP {
			fmt.Printf("%dx%d ", a.b, a.x)
		}
	}

	fmt.Printf("\n")
}
