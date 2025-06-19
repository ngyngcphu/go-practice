# Tic-Tac-Toe

A well-structured command-line tic-tac-toe game implemented in Go, featuring clean architecture with strategy patterns and state management.

## Features

- Interactive command-line gameplay
- Flexible board size support (defaults to 3x3)
- Strategy pattern for different player types
- State pattern for game flow management
- Comprehensive test coverage

## How to Play

1. Players take turns entering their moves
2. Enter row and column coordinates (0-indexed)
3. Example input: `1 2` for row 1, column 2
4. First player to get three in a row wins!

## Project Structure

```
├── cmd/tictactoe/     # Application entry point
├── internal/
│   ├── game/          # Core game logic (Board, Player, Game)
│   ├── states/        # Game state management
│   └── strategy/      # Player strategy implementations
├── pkg/types/         # Shared types and constants
└── main.go           # Main executable
```

## Architecture Highlights

- **Strategy Pattern**: Extensible player strategies (currently supports human players)
- **State Pattern**: Clean game flow management with distinct states
- **Interface Segregation**: Well-defined interfaces for loose coupling
- **Comprehensive Testing**: Unit tests for all major components
