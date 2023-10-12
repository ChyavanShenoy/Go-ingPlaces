Go enforces strong and static typing, meaning variables can only have a single type. A `string` variable that stores "hello world" can not be changed to an `int`, such as the number 3.

One of the biggest benifits of strong typing is that errors can be caught at "compile time". In other words, bugs are more easily caught ahead of time beacuse they are detected when the code is compled before the program is run.

Contrast this with mostr interpreted languages, such as Python, where errors are caught at "runtime". In other words, bugs are detected when the program is run. This can be a problem if the program is run in production, where it is being used by real users.

# String concatenation

Two strings can be combined using the `+` operator. This is called "string concatenation".

# Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello " + "world")
}
```
