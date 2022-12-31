# error_handler

error_handler is an attempt to make debugging easier for programmers and non-programming techies. You get Python or Laravel-like error reporting, with code tracing and easy-to-make-sense-of error report phrases.

## Installation

In your project import the public package:

```bash
go get github.com/joshuaetim/error_handler
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/joshuaetim/error_handler"
)

func main() {
	defer error_handler.HandleError()

	num := 0
	main := 4 / num
	fmt.Println(main)
}
```
Result:
```bash
there was an error:
        File:   /go/src/calc/divisor.go
        Function: main.main
        Line: 13
        Message: You cannot divide by zero
====================================================================================================
8  
9  func main() {
10      defer error_handler.HandleError()
11  
12      num := 0
13      main := 4 / num
        ~~~~~~~~~~~~~~~
14      fmt.Println(main)
15  }
```

## Get in touch

LinkedIn: [https://www.linkedin.com/in/joshuaetim/](https://www.linkedin.com/in/joshuaetim/)

Twitter: [https://twitter.com/theJoshuaEtim](https://twitter.com/theJoshuaEtim)

## License

[Unlicense](https://choosealicense.com/licenses/unlicense/)