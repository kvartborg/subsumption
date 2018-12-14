package subsumption

// Behavior takes a value and a next handler
type Behavior func(value int, next func())

// Hierarchy of behaviors
type Hierarchy struct {
	behaviors []Behavior
	current   int
}

// NewHierarchy returns a new hierarchy of the given behaviors
func NewHierarchy(behaviors ...Behavior) *Hierarchy {
	return &Hierarchy{behaviors, 0}
}

// Behave takes a value which should be passed through the behavior hierarchy
func (h *Hierarchy) Behave(value int) {
	h.current = 0
	h.execute(value)
}

// execute handles the execution of the actual behaviors
func (h *Hierarchy) execute(value int) {
	if h.current >= len(h.behaviors) {
		return
	}

	behavior := h.behaviors[h.current]
	behavior(value, h.nextBehavior(value))
}

// nextBehavior returns a function which will invoke the next behavior if called
// the next method can only be used once within the assigned behavior.
func (h *Hierarchy) nextBehavior(value int) func() {
	called := false
	return func() {
		if called {
			return
		}

		h.current++
		called = true
		h.execute(value)
	}
}
