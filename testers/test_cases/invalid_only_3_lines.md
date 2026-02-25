# Test Case: Invalid — Only 3 Lines Instead of 4

## Description
Rejects tetromino blocks that do not have exactly 4 lines.

## Input (only 3 lines)
```txt
###.
#...
....
```

## Expected Behavior
- Detects block line count != 4.
- Exits with error.

## Expected Output
```txt
ERROR
```

## Purpose
Enforces strict 4-line tetromino blocks.
