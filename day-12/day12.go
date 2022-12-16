package day12

import (
	"sort"
	"strings"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	grid := parseGrid(input)
	start := findNodes(grid, "S")[0]
	end := findNodes(grid, "E")[0]

	steps := bfs(grid, start, end)

	return util.NewSolution(1, steps)
}

func Part2(input string) *util.Solution {
	grid := parseGrid(input)
	start := findNodes(grid, "S")[0]
	end := findNodes(grid, "E")[0]
	startingNodes := findNodes(grid, "a")

	startingNodes = append(startingNodes, start)

	steps := []int{}

	for _, node := range startingNodes {
		stepCount := bfs(grid, node, end)

		if stepCount == notFound {
			continue
		}

		steps = append(steps, stepCount)
	}

	sort.Ints(steps)

	return util.NewSolution(2, steps[0])
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const notFound = -1

type Node struct {
	X      int
	Y      int
	Height int
	Label  string
}

func NewNode(x int, y int, label string) *Node {
	node := &Node{
		X:      x,
		Y:      y,
		Label:  label,
		Height: strings.Index(alphabet, label),
	}

	if label == "S" {
		node.Height = 1
	}

	if label == "E" {
		node.Height = len(alphabet) - 1
	}

	return node
}

func bfs(grid [][]*Node, start *Node, end *Node) int {
	nodesToCheck := queue.New()
	visitedNodes := set.New()

	nodesToCheck.Enqueue(start)
	nodesToCheck.Enqueue(nil)
	visitedNodes.Insert(start)

	depth := 0

	for nodesToCheck.Len() > 0 {
		currentNode := nodesToCheck.Dequeue()

		if currentNode == nil {
			depth++
			nodesToCheck.Enqueue(nil)

			if nodesToCheck.Peek() == nil {
				break
			}

			continue
		}

		if currentNode == end {
			return depth
		}

		adjacentNodes := make([]*Node, 0)

		if currentNode.(*Node).X > 0 {
			adjacentNodes = append(adjacentNodes, grid[currentNode.(*Node).Y][currentNode.(*Node).X-1])
		}

		if currentNode.(*Node).X < len(grid[0])-1 {
			adjacentNodes = append(adjacentNodes, grid[currentNode.(*Node).Y][currentNode.(*Node).X+1])
		}

		if currentNode.(*Node).Y > 0 {
			adjacentNodes = append(adjacentNodes, grid[currentNode.(*Node).Y-1][currentNode.(*Node).X])
		}

		if currentNode.(*Node).Y < len(grid)-1 {
			adjacentNodes = append(adjacentNodes, grid[currentNode.(*Node).Y+1][currentNode.(*Node).X])
		}

		for _, adjacentNode := range adjacentNodes {
			heightDiff := adjacentNode.Height - currentNode.(*Node).Height

			if heightDiff > 1 {
				continue
			}

			if visitedNodes.Has(adjacentNode) {
				continue
			}

			visitedNodes.Insert(adjacentNode)
			nodesToCheck.Enqueue(adjacentNode)
		}
	}

	return notFound
}

func parseGrid(input string) [][]*Node {
	lines := strings.Split(input, "\n")

	grid := make([][]*Node, len(lines))

	for y, line := range lines {
		grid[y] = make([]*Node, len(line))

		for x, label := range strings.Split(line, "") {
			grid[y][x] = NewNode(x, y, label)
		}
	}

	return grid
}

func findNodes(grid [][]*Node, label string) []*Node {
	nodes := make([]*Node, 0)

	for _, row := range grid {
		for _, node := range row {
			if node.Label != label {
				continue
			}

			nodes = append(nodes, node)
		}
	}

	return nodes
}
