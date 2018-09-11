package main

type Case struct {
	a, c []Interval
	b Interval
}

func TestMain(t *testing.T) {
	cases := []Case{
		Case{
			a:[]Interval{Interval{1, 3},Interval{6, 9}},
			b:Interval{2,5},
			c:[]Interval{Interval{1,5},Interval{6,9}},
		},
		Case{
			a:[]Interval{Interval{0, 5},Interval{9, 12}}, 
			b:Interval{7, 16},
			c:[]Interval{Interval{0,5},Interval{7,16}}
		},
		Case{
			a:[]Interval{
				Interval{2, 3},
				Interval{5, 5},
				Interval{6, 6},
				Interval{7, 7},
				Interval{8, 11},
			b:Interval{6,13},
			c:[]Interval{Interval{2,3},Interval{5,5},Interval{6,13}},
		},
	}
}
