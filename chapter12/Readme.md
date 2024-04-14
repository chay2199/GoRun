### Reference & Value Semantics

Common uses of pointers:
- some objects can't be copied safely (mutex)
- some objects are too large to copy
- some methods need to change (mutate) the receiver later
- when decoding protocol data into an object (JSON etc)

```
var r Response
err := json.Unmarshal(j, &r)
```
- when using a pointer to signal a "null" object

The value returned by range is always a copy
```
for i, thing := range things {
    // thing is a copy
    ...
}
```

Use the index if you need to mutate the element:
```
for i := range things {
    things[i].which = whatever
    ...
}
```