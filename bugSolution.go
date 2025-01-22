```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    switch j := i.(type) {
    case int32:
        fmt.Println(j)
    case int:
        fmt.Println(int32(j))
    default:
        fmt.Printf("Unexpected type: %T
", j)
    }
}
```