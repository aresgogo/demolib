// Generated by: main
// TypeWriter: stringer
// Directive: +gen on Pill

package libgogen

import (
	"fmt"
)

const _Pill_name = "PlaceboAspirinIbuprofenParacetamolAcetaminophen"

var _Pill_index = [...]uint8{0, 7, 14, 23, 34, 47}

func (i Pill) String() string {
	if i < 0 || i+1 >= Pill(len(_Pill_index)) {
		return fmt.Sprintf("Pill(%d)", i)
	}
	return _Pill_name[_Pill_index[i]:_Pill_index[i+1]]
}
