# Test Case: Invalid — Wrong Line Length

## Description
Rejects lines that are not exactly 4 characters long.

## Input (first line has 5 chars)
```txt
###..
#...
....
....
```

## Expected Behavior
- Detects a line length != 4.
- Exits with error.

## Expected Output
```txt
ERROR
```

## Purpose
Enforces strict 4x4 format.
