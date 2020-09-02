package libgogen

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestMyTypeSliceWhere(t *testing.T) {
	var mts MyTypeSlice
	mts = []MyType{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}}
	fn := func(mt MyType) bool {
		if math.Sqrt(mt.X*mt.X+mt.Y*mt.Y) < 8 {
			return true
		}
		return false
	}
	got := mts.Where(fn)
	want := mts[:5]
	if !reflect.DeepEqual(got, want) {
		t.Errorf("mts.Where(fn): %#v, want: %#v", got, want)
	}
}

func TestMyTypeSliceCount(t *testing.T) {
	var mts MyTypeSlice
	mts = []MyType{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}}
	fn := func(mt MyType) bool {
		return true
	}
	got := mts.Count(fn)
	want := 9
	if got != want {
		t.Errorf("mts.Count(fn): %d, want: %d", got, want)
	}
}

func TestMyTypeSliceGroupByString(t *testing.T) {
	var mts MyTypeSlice
	mts = []MyType{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}}
	fn := func(my MyType) string {
		return strconv.FormatFloat(my.X, 'f', -1, 64)
	}
	got := mts.GroupByString(fn)
	want := make(map[string]MyTypeSlice)
	for _, v := range mts {
		k := strconv.FormatFloat(v.X, 'f', -1, 64)
		want[k] = append(want[k], v)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("mts.GroupByString(fn): %#v\n want: %#v", got, want)
	}
}

func TestPillString(t *testing.T) {
	var got, want string
	got = Placebo.String()
	want = "Placebo"
	if got != want {
		t.Errorf("Placebo.String(): %s, want: %s", got, want)
	}
	got = Aspirin.String()
	want = "Aspirin"
	if got != want {
		t.Errorf("Aspirin.String(): %s, want: %s", got, want)
	}
	got = Ibuprofen.String()
	want = "Ibuprofen"
	if got != want {
		t.Errorf("Ibuprofen.String(): %s, want: %s", got, want)
	}
	got = Paracetamol.String()
	want = "Paracetamol"
	if got != want {
		t.Errorf("Paracetamol.String(): %s, want: %s", got, want)
	}
	got = Acetaminophen.String()
	want = "Acetaminophen"
	if got != want {
		t.Errorf("Acetaminophen.String(): %s, want: %s", got, want)
	}
}

func TestAggregateString(t *testing.T) {
	employees := EmployeeSlice{
		{"Alice", "Accounting"},
		{"Bob", "Back Office"},
		{"Carly", "Containers"},
	}
	join := func(state string, e Employee) string {
		if state != "" {
			state += ", "
		}
		return state + e.Name
	}
	got := employees.AggregateString(join)
	want := "Alice, Bob, Carly"
	if got != want {
		t.Errorf("employees.AggregateString(join): %s, want: %s", got, want)
	}

	table := func(state string, e Employee) string {
		if state == "" {
			state += fmt.Sprintf("|%-15s|%-15s|\n", "Name", "Department")
			state += fmt.Sprintf("|%-15s|%-15s|\n", e.Name, e.Department)
		} else {
			state += fmt.Sprintf("|%-15s|%-15s|\n", e.Name, e.Department)
		}
		return state
	}
	tableOutput := employees.AggregateString(table)
	t.Logf("table: \n%s", tableOutput)

}

func TestPersonSliceAll(t *testing.T) {
	gang := PersonSlice{
		{"Alice", true},
		{"Bob", false},
		{"Carly", true},
	}
	here := func(p Person) bool {
		return p.Present
	}
	got := gang.All(here) // => false, Bob didn't make it
	want := false
	if got != want {
		t.Errorf("gang.All(here): %t, want: %t", got, want)
	}

	gang = PersonSlice{
		{"Alice", true},
		{"Bob", true},
		{"Carly", true},
	}
	got = gang.All(here) // => true
	want = true
	if got != want {
		t.Errorf("gang.All(here): %t, want: %t", got, want)
	}
}

func TestPersonSliceAny(t *testing.T) {
	people := PersonSlice{
		{"Bueller", false},
		{"Spicoli", false},
		{"Mr. Hand", false},
	}

	bueller := func(p Person) bool {
		return p.Name == "Bueller"
	}

	got := people.Any(bueller) // => true
	want := true
	if got != want {
		t.Errorf("people.Any(bueller): %t, want:%t", got, want)
	}
}

func TestCelsiusSliceAverage(t *testing.T) {
	temps := CelsiusSlice{15.1, -2, 3.6}
	got, err := temps.Average()
	if nil != err {
		t.Error(err)
	}
	want := Celsius((15.1 - 2 + 3.6) / 3)
	if got != want {
		t.Errorf("temps.Average(): %f, want:%f", got, want)
	}
}

func TestPlayerSliceAverageInt(t *testing.T) {
	players := PlayerSlice{
		{"Alice", 450},
		{"Bob", 100},
		{"Carly", 200},
	}

	points := func(p Player) int {
		return p.Points
	}

	got, err := players.AverageInt(points) // => 250, nil
	if nil != err {
		t.Error(err)
	}
	want := 250
	if got != want {
		t.Errorf("players.AverageInt(points): %d, want: %d", got, want)
	}
}

func TestMonsterSliceCount(t *testing.T) {
	monsters := MonsterSlice{
		{"Alice", false, 0},
		{"Bob", true, 4},
		{"Carly", true, 2},
		{"Dave", false, 2},
	}

	werewolf := func(m Monster) bool {
		return m.Fangs > 0 && m.Furry
	}

	got := monsters.Count(werewolf) // => 2 (Bob & Carly)
	want := 2
	if got != want {
		t.Errorf("monsters.Count(werewolf): %d, want:%d", got, want)
	}
}

func TestHipsterSliceDistinct(t *testing.T) {
	hipsters := HipsterSlice{
		{"Neutral Milk Hotel", true, true},
		{"Neutral Milk Hotel", true, true},
		{"Neutral Milk Hotel", true, true},
		{"Neutral Milk Hotel", true, true},
	}

	got := hipsters.Distinct() // => [{"Neutral Milk Hotel", true, true}]
	want := HipsterSlice{{"Neutral Milk Hotel", true, true}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("hipsters.Distinct(): %#v\n got: %#v", got, want)
	}
}

func TestHipsterSliceDistinctBy(t *testing.T) {
	hipsters := HipsterSlice{
		{"Neutral Milk Hotel", true, true},
		{"Death Cab for Cutie", true, true},
		{"You Probably Haven’t Heard of Them", true, true},
		{"Neutral Milk Hotel", false, true},
	}

	band := func(a Hipster, b Hipster) bool {
		return a.FavoriteBand == b.FavoriteBand
	}

	got := hipsters.DistinctBy(band) // => [{"Neutral Milk Hotel", true, true}, {"Death Cab for Cutie", true, true}, {"You Probably Haven’t Heard of Them", true, true}]
	want := HipsterSlice{{"Neutral Milk Hotel", true, true}, {"Death Cab for Cutie", true, true}, {"You Probably Haven’t Heard of Them", true, true}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("hipsters.DistinctBy(band): %#v\n want:%#v", got, want)
	}
}

func TestCustomerSliceFirst(t *testing.T) {
	customers := CustomerSlice{
		{"Alice", false},
		{"Bob", true},
		{"Carly", true},
	}

	come := func(c Customer) bool {
		return c.Here
	}

	got, err := customers.First(come) // => {"Bob", true}, nil
	if nil != err {
		t.Error(err)
	}
	want := Customer{"Bob", true}
	if got != want {
		t.Errorf("customers.First(come): %#v, want:%#v", got, want)
	}
}

func TestMovieSliceGroupByInt(t *testing.T) {
	movies := MovieSlice{
		{"Independence Day", 1996},
		{"Iron Man", 2008},
		{"Fargo", 1996},
		{"Django Unchained", 2012},
		{"WALL-E", 2008},
	}
	year := func(m Movie) int {
		return m.Year
	}
	got := movies.GroupByInt(year)
	want := map[int]MovieSlice{
		1996: {{"Independence Day", 1996}, {"Fargo", 1996}},
		2008: {{"Iron Man", 2008}, {"WALL-E", 2008}},
		2012: {{"Django Unchained", 2012}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("movies.GroupByInt: %#v\n got:%#v", got, want)
	}
}

func TestPriceSliceMax(t *testing.T) {
	prices := PriceSlice{12.34, 43.21, 23.45}
	got, err := prices.Max() // => 43.21
	if nil != err {
		t.Error(err)
	}
	want := Price(43.21)
	if got != want {
		t.Errorf("prices.Max():%f, want:%f", got, want)
	}
}

func TestMovie1SliceMaxDollars(t *testing.T) {
	movies := Movie1Slice{
		{"Independence Day", 1000000},
		{"Iron Man", 5000000},
		{"Fargo", 3000000},
		{"Django Unchained", 9000000},
		{"WALL-E", 4000000},
	}
	box := func(e Movie1) Dollars {
		return e.BoxOffice
	}
	got, err := movies.MaxDollars(box) // => 9000000
	if nil != err {
		t.Error(err)
	}
	want := Dollars(9000000)
	if got != want {
		t.Errorf("movies.MaxDollars: %d, want:%d", got, want)
	}
}

func TestRectangleSliceMaxBy(t *testing.T) {
	rectangles := RectangleSlice{
		{5, 4},
		{6, 7},
		{2, 3},
	}
	area := func(a, b Rectangle) bool {
		return a.Area() < b.Area()
	}

	got, err := rectangles.MaxBy(area) // => {6, 7}
	if nil != err {
		t.Error(err)
	}
	want := Rectangle{6, 7}
	if got != want {
		t.Errorf("rectangles.MaxBy(ares): %#v, want: %#v", got, want)
	}
}

func TestPriceSliceMin(t *testing.T) {
	prices := PriceSlice{12.34, 43.21, 23.45}

	got, err := prices.Min() // => 12.34
	if nil != err {
		t.Error(err)
	}

	want := Price(12.34)
	if got != want {
		t.Errorf("prices.Min(): %f, want:%f", got, want)
	}
}

func TestMovie1SliceMinDollars(t *testing.T) {
	movies := Movie1Slice{
		{"Independence Day", 1000000},
		{"Iron Man", 5000000},
		{"Fargo", 3000000},
		{"Django Unchained", 9000000},
		{"WALL-E", 4000000},
	}

	box := func(e Movie1) Dollars {
		return e.BoxOffice
	}

	got, err := movies.MinDollars(box) // => 1000000
	if nil != err {
		t.Error(err)
	}
	want := Dollars(1000000)
	if got != want {
		t.Errorf("movies.MinDollars(box): %d, want: %d", got, want)
	}
}

func TestRectangleSliceMinBy(t *testing.T) {
	rectangles := RectangleSlice{
		{5, 4},
		{6, 7},
		{2, 3},
	}

	area := func(a, b Rectangle) bool {
		return a.Area() < b.Area()
	}

	got, err := rectangles.MinBy(area) // => {2, 3}
	if nil != err {
		t.Error(err)
	}
	want := Rectangle{2, 3}
	if got != want {
		t.Errorf("rectangles.MinBy(ares): %#v, want:%#v", got, want)
	}
}

func TestPlayerSliceSelectInt(t *testing.T) {
	players := PlayerSlice{
		{"Alice", 450},
		{"Bob", 100},
		{"Carly", 200},
	}

	points := func(p Player) int {
		return p.Points
	}

	got := players.SelectInt(points) // => [450, 100, 200]
	want := []int{450, 100, 200}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("players.SelectInt(points): %#v, want:%#v", got, want)
	}
}

func TestRatingSliceShuffle(t *testing.T) {
	ratings := RatingSlice{1, 2, 3, 4, 5, 6}

	got := ratings.Shuffle() // => {3, 6, 1, 2, 4, 5}
	t.Logf("got: %#v", got)
}

func TestTatingSliceSort(t *testing.T) {
	ratings := RatingSlice{5, 7, 2, 1, 9, 2}

	got := ratings.Sort() // => {1, 2, 2, 5, 7, 9}
	want := RatingSlice{1, 2, 2, 5, 7, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ratings.Sort(): %#v, want: %#v", got, want)
	}

	got = ratings.SortDesc() // => {9, 7, 5, 2, 2, 1}
	want = RatingSlice{9, 7, 5, 2, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ratings.SortDesc(): %#v, want: %#v", got, want)
	}
}

func TestMovieSortBy(t *testing.T) {
	movies := MovieSlice{
		{"Independence Day", 1996},
		{"Iron Man", 2008},
		{"Fargo", 1996},
		{"Django Unchained", 2012},
		{"WALL-E", 2008},
	}

	yearThenTitle := func(a, b Movie) bool {
		if a.Year == b.Year {
			return a.Title < b.Title
		}
		return a.Year < b.Year
	}

	got := movies.SortBy(yearThenTitle)
	want := MovieSlice{
		{"Fargo", 1996},
		{"Independence Day", 1996},
		{"Iron Man", 2008},
		{"WALL-E", 2008},
		{"Django Unchained", 2012},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("movies.SortBy(yearThenTitle): %#v, want: %#v", got, want)
	}
}

func TestBookSliceWhere(t *testing.T) {
	books := BookSlice{
		{"那些忧伤的年轻人", "XuZhiyuan"},
		{"转折年代", "XuZhiyuan"},
		{"纳斯达克的一代", "XuZhiyuan"},
		{"昨日与明日", "XuZhiyuan"},
		{"看见", "ChaiJing"},
		{"用我一辈子去忘记", "ChaiJing"},
	}
	xzy := func(book Book) bool {
		return book.Author == "XuZhiyuan"
	}
	got := books.Where(xzy)
	want := BookSlice{
		{"那些忧伤的年轻人", "XuZhiyuan"},
		{"转折年代", "XuZhiyuan"},
		{"纳斯达克的一代", "XuZhiyuan"},
		{"昨日与明日", "XuZhiyuan"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("books.Where(xzy): %#v, want: %#v", got, want)
	}
}
