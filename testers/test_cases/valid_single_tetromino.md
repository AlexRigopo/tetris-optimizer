# Test Case: Valid Single Tetromino

## Description
Verifies that the program correctly parses a single valid tetromino (4x4 grid, only `#` and `.`, exactly 4 `#`, and a valid connected shape).

## Input
```txt
###.
#...
....
....
```

## Expected Behavior
- Reads exactly 4 lines for the tetromino.
- Each line contains exactly 4 characters.
- Only `#` and `.` characters are present.
- Exactly 4 `#` characters exist.
- Shape is **connected** (no isolated `#` groups).
- Converts `#` → `A` for the first tetromino.

## Expected Output (After Conversion)
```txt
AAA.
A...
....
....
```

## Purpose
Covers the happy path for:
- file parsing
- format validation
- connectivity validation
- letter replacement
