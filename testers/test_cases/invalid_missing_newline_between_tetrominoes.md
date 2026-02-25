# Test Case: Invalid — Missing Newline Between Tetrominoes

## Description
Rejects input where two tetrominoes are not properly separated (expected a blank line between blocks).

## Input (8 consecutive lines, no blank line separator)
```txt
###.
#...
....
....
##..
.#..
....
....
```

## Expected Behavior
- Parser should not treat this as one 8-line tetromino.
- Should detect invalid block formatting / missing separator.
- Exits with error.

## Expected Output
```txt
ERROR
```

## Purpose
Validates that block separation rules are enforced.
