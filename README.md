# thunk
Just an experiment at this stage that may find its way into a real library.

A [typewriter](https://github.com/clipperhouse/typewriter) for use with the golang code generator [gen](https://github.com/clipperhouse/gen).

## Reasoning
Despite golang's claim to being ReallySimple™ to the point of nausea there is a lot of hidden complexity in the interaction between Channels and Goroutines. It's very easy to mess this up and leak resources or deadlock. They're too low level of a primitive to work with directly. Go's lack of support for generics makes this complexity impossible to abstract away (runtime type reflection not withstanding), so we're left with code generation.

## Example

```go
//go:generate gen
package main

import (
	"fmt"
	"time"
)

// +gen thunk:"UnderlyingType"
type String string

func foo() string {
	time.Sleep(2 * time.Second)

	return "Foo!"
}

func main() {
	result := NewStringThunk(foo).Timeout(1 * time.Second).Force()

	fmt.Println("Result is", result)
}
```
