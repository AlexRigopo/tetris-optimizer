# Test Case: Valid Multiple Tetrominoes

## Description
Verifies parsing and conversion of multiple tetrominoes separated by a blank line.

## Input
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
- Splits input into tetromino blocks separated by a single blank line.
- Each block is 4 lines of 4 characters.
- Each block has exactly 4 `#` and is connected.
- First tetromino converts to `A`, second converts to `B`.

## Expected Output (After Conversion)
```txt
AAA.
A...
....
....

BB..
.B..
....
....
```

## Purpose
Ensures:
- multiple-block parsing
- correct separator handling
- sequential letter assignment
