package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	// gridSize defines the width and height of the game grid.
	gridSize = 25

	// liveCell is the character displayed for live cells.
	liveCell = "X"

	// deadCell is the character displayed for dead cells.
	deadCell = "."

	// delay is the duration between generation updates.
	delay = 200 * time.Millisecond

	// generations specifies the total number of generations to simulate.
	generations = 1_000
)

// Grid represents the game's universe as a 2-dimensional boolean array.
// true indicates a live cell, false indicates a dead cell.
type Grid [gridSize][gridSize]bool

// Game encapsulates the current state and generation count for Conway's Game of Life.
type Game struct {
	grid Grid
	gen  int
}

// NewGame creates and initializes a new Game instance with a predefined glider pattern
// positioned near the center of the grid.
//
// Returns:
//   - A pointer to the initialized Game struct.
func NewGame() *Game {
	g := &Game{grid: Grid{}}
	center := gridSize / 2

	g.grid[center][center+1] = true
	g.grid[center+1][center+2] = true
	g.grid[center+2][center] = true
	g.grid[center+2][center+1] = true
	g.grid[center+2][center+2] = true

	return g
}

// ClearScreen clears the terminal screen using ANSI escape codes.
//
// Note: Compatible with most modern terminals.
func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

// LiveNeighbors calculates the number of live neighbors around a specific cell.
// It uses toroidal wrapping at grid edges.
//
// Parameters:
//   - x: The X-coordinate (row index) of the target cell.
//   - y: The Y-coordinate (column index) of the target cell.
//
// Returns:
//   - The count of live neighboring cells (0-8).
func (g *Game) LiveNeighbors(x, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := (x+i+gridSize)%gridSize, (y+j+gridSize)%gridSize
			if g.grid[nx][ny] {
				count++
			}
		}
	}

	return count
}

// NextGen computes the next generation by applying the rules of Conway's Game of Life
// to each cell, updating the game's internal grid state.
//
// Rules applied:
//   - Any live cell with 2 or 3 neighbors survives.
//   - Any dead cell with exactly 3 neighbors becomes alive.
//   - All other cells die or remain dead.
func (g *Game) NextGen() {
	var newGrid Grid

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			neighbors := g.LiveNeighbors(x, y)
			newGrid[x][y] = neighbors == 3 || (g.grid[x][y] && neighbors == 2)
		}
	}

	g.grid = newGrid
	g.gen++
}

// CountLiveCells counts the total number of currently live cells on the grid.
//
// Returns:
//   - The count of live cells as an integer.
func (g *Game) CountLiveCells() int {
	count := 0

	for x := range g.grid {
		for y := range g.grid[x] {
			if g.grid[x][y] {
				count++
			}
		}
	}

	return count
}

// Print outputs the current state of the grid to the terminal with clear visual borders.
func (g *Game) Print() {
	border := strings.Repeat("─", gridSize+2)
	fmt.Println("┌" + border + "┐")

	for _, row := range g.grid {
		fmt.Print("│ ")
		for _, cell := range row {
			if cell {
				fmt.Print(liveCell)
			} else {
				fmt.Print(deadCell)
			}
		}
		fmt.Println(" │")
	}

	fmt.Println("└" + border + "┘")
}

// main starts and runs the simulation for a predefined number of generations.
// It initializes the game, updates the grid state, and prints each generation.
//
// Simulation can be stopped manually by pressing Ctrl+C.
func main() {
	game := NewGame()

	for i := 0; i < generations; i++ {
		ClearScreen()

		fmt.Printf("Conway's Game of Life - Generation: %d | Live Cells: %d\n",
			game.gen, game.CountLiveCells())

		game.Print()
		fmt.Println("Press Ctrl+C to exit")

		game.NextGen()
		time.Sleep(delay)
	}

	// Final state display
	ClearScreen()
	fmt.Printf("Conway's Game of Life - Final Generation: %d | Live Cells: %d\n",
		game.gen, game.CountLiveCells())
	game.Print()
}
