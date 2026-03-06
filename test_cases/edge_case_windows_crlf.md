# Edge Case: Windows CRLF Line Endings

## Description
Ensures parser accepts files encoded with `\r\n` line endings.

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
- Carriage returns are trimmed from each line.
- Blocks validate successfully.

## Expected Output
- Valid solution printed.

## Purpose
Prevents platform-specific failures on Windows-authored files.
