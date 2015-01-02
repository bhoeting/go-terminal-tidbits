#terminal-tidbits

Some simple terminal utilities for Go.

##Includes
* Colored output

##Install
`go get github.com/bhoeting/terminal-tidbits/`

##Use
```go
import (
  "github.com/bhoeting/terminal-tidbits/"
)
```

###Colors
```go
// Write to console
term.Color("@blue(Hello, world!)")
term.Colorf("@green(%s), @yellow(%s)%s\n", "Hello", "world", "!")

// Write to string
var str1 string = term.SColor("@blue(Hello, world!)")
var str2 string = term.SColorf("@green(%s), @yellow(%s)%s\n", "Hello", "world", "!")
```

