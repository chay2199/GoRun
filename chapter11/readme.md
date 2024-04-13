### Structs, struct tags & JSON

Two struct types are compatible if:
- the fields have the same types and names
- in the same order
- and the same tags (*)

A struct may be copied or passed as a parameter in its entirety

A stuct is comparable if all its fields are comparable

The zero value for a struct is "zero" for each field in turn