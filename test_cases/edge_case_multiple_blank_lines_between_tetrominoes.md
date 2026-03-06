# Edge Case: Multiple Blank Lines Between Tetrominoes

## Description
Checks parser behavior when valid blocks are separated by more than one blank line.

## Input
```txt
###.
#...
....
....


##..
##..
....
....
```

## Expected Behavior
- Parser ignores empty separators.
- Both tetrominoes are loaded and validated.

## Expected Output
- Valid solution printed.

## Purpose
Documents behavior for extra separator lines in user input files.
