## help: Print this help message
.PHONY: help
help:
	@echo 'Usage':
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run/game: Run the game.
.PHONY: run/game
run/game:
	go run ./cmd