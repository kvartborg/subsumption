package subsumption_test

import (
	"testing"

	"github.com/kvartborg/subsumption"
	"github.com/kvartborg/subsumption/behavior"
)

func TestIndexOutOfRange(t *testing.T) {
	behavior := func(value int, next func()) { next() }
	behaviors := subsumption.NewHierarchy(behavior, behavior)
	behaviors.Behave(1)
}

func Example() {
	behaviors := subsumption.NewHierarchy(
		behavior.Avoid,
		behavior.Wander,
		behavior.Explore,
	)

	for i := 20; i > 0; i-- {
		behaviors.Behave(i)
	}

	// Output: Exploring
	// Exploring
	// Exploring
	// Exploring
	// Exploring
	// Exploring
	// Exploring
	// Exploring
	// Exploring
	// Exploring
	// Wandering
	// Wandering
	// Wandering
	// Wandering
	// Avoiding
	// Avoiding
	// Avoiding
	// Avoiding
	// Avoiding
	// Avoiding
}
