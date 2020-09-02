// Generated by: main
// TypeWriter: slice
// Directive: +gen on Customer

package libgogen

import "errors"

// CustomerSlice is a slice of type Customer. Use it where you would use []Customer.
type CustomerSlice []Customer

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv CustomerSlice) First(fn func(Customer) bool) (result Customer, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no CustomerSlice elements return true for passed func")
	return
}
