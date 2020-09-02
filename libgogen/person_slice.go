// Generated by: main
// TypeWriter: slice
// Directive: +gen on Person

package libgogen

// PersonSlice is a slice of type Person. Use it where you would use []Person.
type PersonSlice []Person

// All verifies that all elements of PersonSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv PersonSlice) All(fn func(Person) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Any verifies that one or more elements of PersonSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv PersonSlice) Any(fn func(Person) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}