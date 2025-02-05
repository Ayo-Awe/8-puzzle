package main

import (
	"fmt"
	"slices"
)

// x,y
var directions = map[string][2]int{
	"<": {-1, 0},
	">": {1, 0},
	"v": {0, 1},
	"^": {0, -1},
}

var target = [3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", " "},
}

var X, Y = 0, 1

func main() {
	puzzle := [3][3]string{
		{"8", "1", "3"},
		{"4", " ", "2"},
		{"7", "6", "5"},
	}
	fmt.Println(solve(puzzle, target))
}

func solve(puzzle, target [3][3]string) []string {
	type state struct {
		grid [3][3]string
		path []string
		pos  [2]int
	}

	// find start position
	start := findStart(puzzle, " ")

	queue := []state{}
	seen := make(map[[3][3]string]int)

	initial := state{grid: puzzle, pos: start, path: nil}
	queue = append(queue, initial)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// we've seen this position before, but with a shorter path
		if steps, ok := seen[curr.grid]; ok && steps < len(curr.path) {
			continue
		}

		// check if we've found a solution
		if curr.grid == target {
			return curr.path
		}

		// explore all positions from this state
		for dir, coord := range directions {
			pos := [2]int{curr.pos[X] + coord[X], curr.pos[Y] + coord[Y]}

			if pos[X] < 0 || pos[X] >= 3 || pos[Y] < 0 || pos[Y] >= 3 {
				continue
			}

			// update newGrid
			grid := curr.grid
			grid[pos[Y]][pos[X]], grid[curr.pos[Y]][curr.pos[X]] = grid[curr.pos[Y]][curr.pos[X]], grid[pos[Y]][pos[X]]

			path := slices.Clone(curr.path)
			path = append(path, dir)

			// add state to queue
			queue = append(queue, state{path: path, grid: grid, pos: pos})
		}

		// mark as seen
		seen[curr.grid] = len(curr.path)
	}

	// no solution
	return nil
}

func findStart(puzzle [3][3]string, startSymbol string) [2]int {
	for y := range puzzle {
		for x := range puzzle[y] {
			if puzzle[y][x] == startSymbol {
				return [2]int{x, y}
			}
		}
	}
	return [2]int{}
}
