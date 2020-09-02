// Generated by: main
// TypeWriter: slice
// Directive: +gen on Celsius

package libgogen

import "errors"

// CelsiusSlice is a slice of type Celsius. Use it where you would use []Celsius.
type CelsiusSlice []Celsius

// Average sums CelsiusSlice over all elements and divides by len(CelsiusSlice). See: http://clipperhouse.github.io/gen/#Average
func (rcv CelsiusSlice) Average() (Celsius, error) {
	var result Celsius

	l := len(rcv)
	if l == 0 {
		return result, errors.New("cannot determine Average of zero-length CelsiusSlice")
	}
	for _, v := range rcv {
		result += v
	}
	result = result / Celsius(l)
	return result, nil
}