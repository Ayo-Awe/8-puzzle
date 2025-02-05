# 8-Puzzle

This project implements a simple **8-puzzle solver** in Go. The program attempts to find the shortest sequence of moves needed to transform an initial 3x3 puzzle state into the target solved state.

## Target State

The goal is to transform any given puzzle into the following state:

```
1 2 3
4 5 6
7 8
```

(The blank space represents the empty tile.)

## Example Initial State

```
8 1 3
4   2
7 6 5
```

## How It Works

1. The program initializes the puzzle state.
2. It finds the blank space (`" "`) position.
3. It explores all valid moves using **BFS**, keeping track of the visited states.
4. If a solution is found, the sequence of moves (`<`, `>`, `v`, `^`) is returned.
5. If no solution exists, it returns `nil`.

## Running the Program

To run the solver, simply execute:

```sh
 go run main.go
```

## TODO

- Build a UI to play the 8-puzzle game and use solver as AI hints
- Implement random valid state generation
