# thunk
Just an experiment at this stage that may find its way into a real library.

A [typewriter](https://github.com/clipperhouse/typewriter) for use with the golang code generator [gen](https://github.com/clipperhouse/gen).

## Reasoning
Despite golang's claim of being ReallySimple™ there is a lot of hidden complexity in the interaction between Channels and Goroutines. It's very easy to mess this up and leak resources or deadlock. They're too low level of a primitive to reliably work with. Go's lack of support for generics makes this complexity difficult to abstract away (runtime type reflection not withstanding), so we're left with code generation.

## Example

```go
package main

import (
	"fmt"
	"time"
)

// +gen thunk slice:"Where,GroupBy[int]"
type User struct {
	id   int
	name string
	age  int
}

func getUsers() []User {
	time.Sleep(2 * time.Second)
	return UserSlice{
		{1, "Alice", 18},
		{2, "Bob", 18},
		{3, "Carley", 21},
		{4, "David", 16},
	}

}

var isOver18 = func(u User) bool { return u.age >= 18 }
var age = func(u User) int { return u.age }

func main() {
	result := UsersThunk(getUsers).
		Timeout(1 * time.Second).
		Force().
		Where(isOver18).
		GroupByInt(age).

	fmt.Println("Result is", result)
}

```
