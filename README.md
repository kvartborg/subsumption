# Subsumption

This is a very simple implementation of the subsumption architecture, it receives
a value and passes it through the behavior hierarchy.

```go
package main

import (
    "github.com/kvartborg/subsumption"
    "github.com/kvartborg/subsumption/behavior"
)

func main() {
    behaviors := subsumption.NewHierarchy(
        behavior.Avoid,
        behavior.Wander,
        behavior.Explore,
    )

    for i := 20; i > 0; i-- {
        behaviors.Behave(i)
    }
}
```
