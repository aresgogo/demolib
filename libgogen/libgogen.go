package libgogen

// MyType is a go-struct-template for generating
// `MyTypeSlice with Where,Count,GroupBy[string] method` code by the "+gen" Annotation bellow
//
// +gen slice:"Where,Count,GroupBy[string]"
type MyType struct {
	X float64
	Y float64
}

// Pill is a go-int-template for generating
// `Pill.String() method` code by the "+gen" Annotation bellow
//
// +gen stringer
type Pill int

const (
	// Placebo is a harmless pill, medicine, or procedure prescribed more for the psychological benefit to the patient than for any physiological effect.
	Placebo Pill = iota
	// Aspirin is a synthetic compound used medicinally to relieve mild or chronic pain and to reduce fever and inflammation.
	Aspirin
	// Ibuprofen is a synthetic compound used widely as an analgesic and anti-inflammatory drug.
	Ibuprofen
	// Paracetamol is a synthetic compound used as a drug to relieve and reduce fever, usually taken in tablet form; acetaminophen.
	Paracetamol
	// Acetaminophen is an analgesic drug used to treat headaches, arthritis, etc.,
	// and also to reduce fever, often as an alternative to aspirin. Proprietary names include Tylenol.
	Acetaminophen
)

// Employee is a go-struct-template for generating
// `EmployeeSlice with Aggregate[string] method` code by the "+gen" Annotation bellow
//
// +gen slice:"Aggregate[string]"
type Employee struct {
	Name       string
	Department string
}

// Person is a go-struct-template for generating
// `PersonSlice with All,Any method` code by the "+gen" Annotation bellow
//
// +gen slice:"All,Any"
type Person struct {
	Name    string
	Present bool
}

// Celsius is a go-float64-template for generating
// `CelsiusSlice with Average method` code by the "+gen" Annotation bellow
//
// +gen slice:"Average"
type Celsius float64

// Player is a go-struct-template for generating
// `PlayerSlice with Average[int] method` code by the "+gen" Annotation bellow
//
// +gen slice:"Average[int],Select[int]"
type Player struct {
	Name   string
	Points int
}

// Monster is a go-struct-template for generating
// `MonsterSlice with Count method` code by the "+gen" Annotation bellow
//
// +gen slice:"Count"
type Monster struct {
	Name  string
	Furry bool
	Fangs int
}

// Hipster is a go-struct-template for generating
// `HipsterSlice with Distinct,DistinctBy method` code by the "+gen" Annotation bellow
//
// +gen slice:"Distinct,DistinctBy"
type Hipster struct {
	FavoriteBand string
	Mustachioed  bool
	Bepectacled  bool
}

// Customer is a go-struct-template for generating
// `CustomerSlice with First method` code by the "+gen" Annotation bellow
//
// +gen slice:"First"
type Customer struct {
	Name string
	Here bool
}

// Movie is a go-struct-template for generating
// `MovieSlice with GroupBy[int] method` code by the "+gen" Annotation bellow
//
// +gen slice:"GroupBy[int],SortBy"
type Movie struct {
	Title string
	Year  int
}

// Price is go-float64-template for generating
// `PriceSlice with Max,Min method` code by the "+gen" Annotation bellow
//
// +gen slice:"Max,Min"
type Price float64

// Movie1 is a go-struct-template for generating
// `Movie1 with max[Dollars],Min[Dollars] method` code by the "+gen" Annotation bellow
//
// +gen slice:"Max[Dollars],Min[Dollars]"
type Movie1 struct {
	Title     string
	BoxOffice Dollars
}

// Dollars ia a structure based on `int` type
type Dollars int

// Rectangle is a go-struct-template for generating
// `RectangleSlice with MaxBy method` code by the "+gen" Annotation bellow
//
// +gen slice:"MaxBy,MinBy"
type Rectangle struct {
	Width, Height int
}

// Area is a method of Rectangle, used to calculate rectangle's area
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Rating is a go-int-template for generating
// `RatingSlice with Shuffle` code by the "+gen" Annotation bellow
//
// +gen slice:"Shuffle,Sort,SortDesc"
type Rating int

// Book is a go-struct-template for generating
// `BookSlice with Where` code by the "+gen" Annotation bellow
//
// +gen slice:"Where"
type Book struct {
	Name   string
	Author string
}
