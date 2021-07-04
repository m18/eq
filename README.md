# eq — equality checks

Equality checks are available for

- int slices
- byte slices
- string slices
- string maps
- string to simple type maps
- string sets

## Example

```go
package main

import (
	"fmt"

	"github.com/m18/eq"
)

func main() {
	equal := eq.IntSlices([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(equal)

	equal = eq.IntSlices([]int{1, 2, 3}, []int{3, 2, 1})
	fmt.Println(equal)
}
```

Output
```
true
false
```