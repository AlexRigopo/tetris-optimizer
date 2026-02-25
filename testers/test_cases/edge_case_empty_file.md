# Edge Case: Empty File

## Description
Rejects an empty input file.

## Input
(empty file)

## Expected Behavior
- Detects no content / no tetromino blocks.
- Exits with error.

## Expected Output
```txt
ERROR
```

## Purpose
Avoids undefined behavior on empty input.
