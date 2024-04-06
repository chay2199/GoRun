Control Statements; Declarations & Types

Go doesn't have while loops. You can have an infinite loop using:
```
for {

}
```

A switch is another choice between alternatives. It is a shortcut replacing a series of if then statements.

```
switch a := f.Get() a {
    case 0, 1, 2:
        fmt.Println("Possible")
    case 3, 4, 5, 6:
        fmt.Println("Impossible")
    default:
        // default code
}
```

Every name that's capitalized is exported. That means another package in the program can import it.

Within a package, everything is visible even across files.

A package "A" cannot import a package that imports A. Cycles are not allowed.

init function is called before main


```
func BadRead(f *os.File, buf []byte) error {
    var err error
    for {
        n, err := f,Read(buf) // this err is local to the for loop
        if err != nil {
            break
        }
        foo(buf)
    }
    return err // will always be nil
}
```