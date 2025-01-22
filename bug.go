```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    j := i.(int32)
    fmt.Println(j)
}
```