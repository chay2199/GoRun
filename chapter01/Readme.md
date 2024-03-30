fmt.Print() prints whatever it gets to the terminal without adding any space or newlines unless it’s unequivocally coded.

```
package main

import "fmt"

func main() {
    fmt.Print("Hello")
    fmt.Print("World")
}
```

fmt.Printf() provides custom formatting of input strings using one or more verbs and then prints the formatted string to the terminal without appending any space or newlines (unless explicitly coded).

```
package main

import "fmt"

func main() {
    year := 2001

    fmt.Printf("I was born in %d", year)

}
```

%v represents the named value in its default format
%T represents the type of the value
%d expects value to be an integer type of base 10
%b expects value to be an integer type of base 2
%s the bytes of string or slice
%f expects value to have a float type

fmt.Println() is similar to fmt.Print(), the difference being that it adds spaces between arguments and appends a new line at the end.

```
package main

import "fmt"

func main() {
    fmt.Println("Hello", "World")

    fmt.Println("welcome to my world!")

}
```

The fmt.Sprint() function formats and returns the input string without printing anything to the terminal.

```
package main

import "fmt"

func main() {
    a := fmt.Sprint("Hello World")
    b := fmt.Sprint("Hello", "World")

    fmt.Println(a)
    fmt.Println(b)

}
```

fmt.Sprintln() is similar in function to fmt.Sprint(), except that it automatically adds spaces between arguments.

```
package main

import "fmt"

func main() {
    a := fmt.Sprintln("Hello","World")

    fmt.Println(a)

}
```

fmt.Sprintf() is used to format an input string. It also works like fmt.Printf(), the significant distinction being that fmt.Sprintf() returns the value as opposed to printing it out.

```
package main

import "fmt"

func main() {
    name := "Chisom"
    s := fmt.Sprintf("My name is %s", name)

    fmt.Print(s)

}
```

fmt.Scan() collects user input from the standard input and stores this in successive arguments. Spaces or newlines are considered multiple values and are stored in multiple arguments. This function anticipates that an address of each argument should be passed.

```
package main

import "fmt"

func main() {

    var name string
	fmt.Println("What's your name?")
	
    fmt.Scan(&name)
	fmt.Println("Nice to meet you", name)
}
```

fmt.Scanf() reads text from the standard input and puts away progressive space-separated values into progressive arguments as specified by the format. It also expects an address of each argument to be passed. The Scan method of the fmt package will read a single space-separated token (like a single word) from the standard input device (which is usually the console). If more than one word is entered, only the first one is read by Scan. The remainder is left in the console and can be read by another Scan call.

```
package main

import "fmt"

func main() {
    var name string

	fmt.Printf("What's your name: ")

	fmt.Scanf("%s",&name)

	fmt.Printf("Your name is %s\n", name)
}
```

fmt.Scanln() functions similarly to fmt.Scan, but it stops scanning at a newline. This means that after the last item, there must be a newline or EOF. Si ut reads space-separated tokens till EOF

```
package main

import "fmt"

func main() {
  	var  z int

	fmt.Scanln(&z)

	fmt.Printf("Scanln: Z = %d\n", z)
}
```
