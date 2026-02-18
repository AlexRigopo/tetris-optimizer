# Test Case: Invalid — Invalid Character

## Description
Rejects characters other than `#` and `.`.

## Input (contains 'X')
```txt
###.
#..X
....
....
```

## Expected Behavior
- Detects invalid character `X`.
- Exits with error.

## Expected Output
```txt
ERROR
```

## Purpose
Prevents parsing of corrupted input.
