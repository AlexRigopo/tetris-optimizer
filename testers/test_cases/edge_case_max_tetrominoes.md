# Edge Case: Maximum Tetrominoes (A–Z)

## Description
Tests the upper boundary of letter mapping (commonly 26 tetrominoes if using A–Z).

## Input
- 26 valid tetromino blocks (each 4x4, connected, exactly 4 `#`)
- Each block separated by a blank line

## Expected Behavior
- Converts tetrominoes to letters `A` through `Z` without overflow.
- Does not panic or mis-label blocks.
- Assembler attempts to produce the smallest square solution.

## Expected Output
- A valid minimal square layout printed (exact layout depends on your solver).

## Purpose
Validates:
- scaling
- alphabet boundary handling
- memory stability
