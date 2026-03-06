# Test Case: Invalid — Wrong Number of '#'

## Description
Rejects tetrominoes that do not contain exactly 4 `#`.

## Input (5 '#')
```txt
###
...#
....
....
```

## Expected Behavior
- Detects 5 `#` characters in a 4x4 block.
- Exits with error and prints an error message to stderr.

## Expected Output
```txt
ERROR
```

## Purpose
Guards against malformed tetromino definitions.
