### Composition

The fields of an embedded struct are promoted to the level of the embedding structure

```
type Pair struct {
    Path string
    Hash string
}

type PairWithLength {
    Pair
    Length int
}
pl := PairWithLength{Pair{"/usr", "0xfdd3"}, 123}
fmt.Println(pl.Path, pl.Length) // not pl.x.Path
```

The methods of an embedded struct are also promoted. Those methods can't see the fields of the embedding struct.

The embedding structure may declare the same methods and so override the promoted methods.

We can make an interface that will satisfy with a method promoted from the lower struct.

sort.Reverse is defined as:
```
type reverse struct {
    // This embedded Interface permits Reverse to use the methods of another Interface implementation
    Interface
}

func (r reverse) Less(i, j int) bool {
    return r.Interface.Less(j, i)
}

// Reverse returns the reverse order of the data
func Reverse(data Interface) Interface {
    return &reverse{data}
}
```