### FUNCTIONS, PARAMETERS AND DEFER

PARAMETERS THAT ARE PASSED BY VALUE:
- numbers
- bool
- arrays
- structs

BY REFERENCE:
- things passed by pointer (&x)
- strings (but they're immutable)
- slices
- maps
- channels

Unlike a closure, defer copies arguments to the deferred call.

A defer statement runs before the return is done