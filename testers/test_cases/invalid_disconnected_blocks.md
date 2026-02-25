# Test Case: Invalid — Disconnected '#' Blocks (Invalid Shape)

## Description
Rejects tetrominoes where the 4 `#` cells are not orthogonally connected as a single shape.

## Input (4 '#' but disconnected)
```txt
#...
....
..#.
..#.
```

## Expected Behavior
- Counts exactly 4 `#`, but detects that the shape is **not connected**.
- Exits with error.

## Expected Output
```txt
ERROR
```

## Purpose
Ensures connectivity validation (4-neighbor adjacency) is implemented.
