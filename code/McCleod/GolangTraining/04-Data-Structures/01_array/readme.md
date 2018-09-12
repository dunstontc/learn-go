# Array

- definition
  - An array is a numbered sequence of elements of a single type.
  - The number of elements is called the length and is never negative. 
  - The length is part of the array's type; it must evaluate to a non-negative constant representable by a value of type int. 
  - The length of an array a can be discovered using the built-in function len. 
  - The elements can be addressed by integer indices 0 through len(a)-1. 
  - Array types are always one-dimensional but may be composed to form multi-dimensional types. 
  - not dynamic
    - does not change in size
- a basic array
  - len
  - index access
  - assigning a value to an index position in an array

## Array Examples
- understanding the difference between index position and the items stored
  - if you're storing three items in array a, those items will be at index positions 0, 1, 2
    - len(a)-1 is your last index position
      - eg, 3-1 = 2 → 2 is your last index position for your array, a, which has three items
- using break in a loop
