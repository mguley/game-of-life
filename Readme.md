# Conway's Game of Life in Go

[![Go Version](https://img.shields.io/badge/go-1.24-blue)](https://golang.org/dl/)

This repository contains a simple and performant implementation of **Conway's Game of Life** written in Go.

## Quick Start

To run the simulation locally:

```sh
make run/game
```

Or directly via Go:

```sh
go run ./cmd
```

## How It Works

- `Universe Size`: 25x25 grid (configurable in code)
- `Initial Pattern`: Classic "Glider" pattern initialized at the center
- `Visualization`: Terminal-based with live (`X`) and dead (`.`) cells
- `Generations`: Runs for 1_000 generations or until manually terminated

### Demo

![Conway's Game of Life Demo](docs/demo.gif)

## Project Structure
```
game-of-life/
├── cmd/
│   └── main.go       # Main entry point
├── docs/
│   └── demo.gif      # GIF animation of the app
├── Makefile          # Quick commands
├── go.mod            # Go module file
└── Readme.md         # Project overview and instructions
```

## Available Make Commands

- Run the game
```sh
make run/game
```

- Help
```sh
make help
```

## Future Improvements & Ideas
The following enhancements are planned for future releases:
- `Concurrent Generation Calculation`: Use Go goroutines to parallelize calculations, boosting performance for larger grids.
- `Dynamic Patterns`: Allow custom initial patterns from file inputs (JSON or plain-text).
- `Interactive Controls`: Add pause/resume functionality and speed controls.
- `Configurable Parameters`: Provide command-line flags for grid size, generation count, simulation speed, and initial patterns.
- `Improved Rendering`: Reduce terminal flickering or integrate advanced rendering libraries.
