#terminal-tidbits

Some simple terminal utilities for Go.

##Includes
* ✔ Colored output
* Column formatting
* Table display

##Install
`go get github.com/bhoeting/terminal-colors/`

##Use
```go
import (
  "github.com/bhoeting/terminal-colors/"
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

// Escaping @
term.Color("@blue(brennan.hoeting)(@)@green(gmail.com)")
```

