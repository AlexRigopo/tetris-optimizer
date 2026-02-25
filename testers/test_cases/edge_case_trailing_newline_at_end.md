# Edge Case: Trailing Newline at End of File

## Description
Accepts valid input even if the file ends with an extra newline after the last tetromino.

## Input (note the extra blank line at end)
```txt
###.
#...
....
....

```

## Expected Behavior
- Should parse the single tetromino successfully.
- Should **not** create an extra empty block due to the trailing newline.
- Converts `#` → `A`.

## Expected Output (After Conversion)
```txt
AAA.
A...
....
....
```

## Purpose
Prevents false errors caused by common file formatting (final newline).
