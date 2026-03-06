# Test Case: Invalid - Space Character In Grid

## Description
Rejects tetromino lines that contain spaces instead of `.`.

## Input
```txt
##..
# ..
#...
....
```

## Expected Behavior
- Detects invalid character (space).
- Exits with error.

## Expected Output
```txt
ERROR
```

## Purpose
Catches malformed inputs that look visually aligned but violate format rules.
