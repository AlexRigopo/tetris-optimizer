# Tetris Optimizer (Go)

Tetris Optimizer reads a list of tetrominoes from a text file, **labels each piece** (`A`, `B`, `C`, …), and then packs them into the **smallest possible square** without overlaps.

The goal is the classic “fit these tetrominoes into the minimum board size” problem: start from the smallest square that could possibly contain all blocks, try to place every piece, and grow the square only if needed.

---

## What this project does

- Reads tetrominoes from a file (4×4 blocks of `#` and `.`)
- Validates the input format and tetromino correctness
- Converts each tetromino’s `#` into a unique uppercase letter (`A`, `B`, `C`, …)
- Solves the packing problem and prints the resulting square
- Includes unit tests and ready-made test cases

---

## Project layout

```
.
├── main.go
├── go.mod
├── example.txt
│
├── actions/
│   ├── loadtetrominoes.go
│   ├── validatetetrominoes.go
│   ├── labeltetrominoes.go
│   └── solvesmalestsquare.go
│
└── testers/
    ├── loadtetrominoes_test.go
    ├── validatetetrominoes_test.go
    ├── labeltetrominoes_test.go
    ├── solvesmallestsquare_test.go
    ├── valid_tetrominoes.txt
    └── test_cases/
```

---

## Input format

Each tetromino is described as **4 lines**, and each line must be **exactly 4 characters**.

- Allowed characters: `#` (block) and `.` (empty)
- Each tetromino must contain **exactly 4 `#` blocks**
- Blocks must be **connected** (no isolated `#`)
- Multiple tetrominoes are separated by a **blank line**

### ✅ Valid example

```
###.
#...
....
....
```

### ❌ Invalid example (5 blocks)

```
###.
##..
....
....
```

---

## Run the program

From the project root:

```bash
go run . <filename>
```

Example:

```bash
go run . example.txt
```

If the argument count is wrong:

```
Usage: go run . <filename>
```

If the filename is empty:

```
Error: No file name
```

---

## How the solver works (high level)

1. **Load** tetrominoes from the file.
2. **Validate** each tetromino’s shape and formatting.
3. **Label** each piece (`A`, `B`, `C`, …) so the final output is readable.
4. Compute a starting square size (based on total blocks).
5. Use a **backtracking** placement algorithm:
   - Try to place the next tetromino in every possible spot.
   - If stuck, undo (backtrack) and try the next position.
6. If no solution exists for the current size, **increase the square** and retry.

This guarantees the **smallest** square that can fit all given pieces.

---

## Run tests

```bash
go test ./...
```

Test inputs live in `testers/test_cases/` (plus `testers/valid_tetrominoes.txt`).

---

## Requirements

- Go (any modern version compatible with `go.mod`)
- No external dependencies
