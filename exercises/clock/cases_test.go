package clock

// Source: exercism/problem-specifications
// Commit: b344762 clock: Add test case for exactly negative sixty minutes.
// Problem Specifications Version: 2.4.0

// Create a new clock with an initial time
var timeTests = []struct {
	n    string
	h, m int
	want string
}{
	{
		"on the hour",
		8, 0, "08:00",
	},
	{
		"past the hour",
		11, 9, "11:09",
	},
	{
		"midnight is zero hours",
		24, 0, "00:00",
	},
	{
		"hour rolls over",
		25, 0, "01:00",
	},
	{
		"hour rolls over continuously",
		100, 0, "04:00",
	},
	{
		"sixty minutes is next hour",
		1, 60, "02:00",
	},
	{
		"minutes roll over",
		0, 160, "02:40",
	},
	{
		"minutes roll over continuously",
		0, 1723, "04:43",
	},
	{
		"hour and minutes roll over",
		25, 160, "03:40",
	},
	{
		"hour and minutes roll over continuously",
		201, 3001, "11:01",
	},
	{
		"hour and minutes roll over to exactly midnight",
		72, 8640, "00:00",
	},
	{
		"negative hour",
		-1, 15, "23:15",
	},
	{
		"negative hour rolls over",
		-25, 0, "23:00",
	},
	{
		"negative hour rolls over continuously",
		-91, 0, "05:00",
	},
	{
		"negative minutes",
		1, -40, "00:20",
	},
	{
		"negative minutes roll over",
		1, -160, "22:20",
	},
	{
		"negative minutes roll over continuously",
		1, -4820, "16:40",
	},
	{
		"negative sixty minutes is previous hour",
		2, -60, "01:00",
	},
	{
		"negative hour and minutes both roll over",
		-25, -160, "20:20",
	},
	{
		"negative hour and minutes both roll over continuously",
		-121, -5810, "22:10",
	},
}

// Add minutes
var addTests = []struct {
	n       string
	h, m, a int
	want    string
}{
	{
		"add minutes",
		10, 0, 3, "10:03",
	},
	{
		"add no minutes",
		6, 41, 0, "06:41",
	},
	{
		"add to next hour",
		0, 45, 40, "01:25",
	},
	{
		"add more than one hour",
		10, 0, 61, "11:01",
	},
	{
		"add more than two hours with carry",
		0, 45, 160, "03:25",
	},
	{
		"add across midnight",
		23, 59, 2, "00:01",
	},
	{
		"add more than one day (1500 min = 25 hrs)",
		5, 32, 1500, "06:32",
	},
	{
		"add more than two days",
		1, 1, 3500, "11:21",
	},
}

// Subtract minutes
var subtractTests = []struct {
	n       string
	h, m, a int
	want    string
}{
	{
		"subtract minutes",
		10, 3, 3, "10:00",
	},
	{
		"subtract to previous hour",
		10, 3, 30, "09:33",
	},
	{
		"subtract more than an hour",
		10, 3, 70, "08:53",
	},
	{
		"subtract across midnight",
		0, 3, 4, "23:59",
	},
	{
		"subtract more than two hours",
		0, 0, 160, "21:20",
	},
	{
		"subtract more than two hours with borrow",
		6, 15, 160, "03:35",
	},
	{
		"subtract more than one day (1500 min = 25 hrs)",
		5, 32, 1500, "04:32",
	},
	{
		"subtract more than two days",
		2, 20, 3000, "00:20",
	},
}

// Compare two clocks for equality
type hm struct{ h, m int }

var eqTests = []struct {
	n      string
	c1, c2 hm
	want   bool
}{
	{
		"clocks with same time",
		hm{15, 37},
		hm{15, 37},
		true,
	},
	{
		"clocks a minute apart",
		hm{15, 36},
		hm{15, 37},
		false,
	},
	{
		"clocks an hour apart",
		hm{14, 37},
		hm{15, 37},
		false,
	},
	{
		"clocks with hour overflow",
		hm{10, 37},
		hm{34, 37},
		true,
	},
	{
		"clocks with hour overflow by several days",
		hm{3, 11},
		hm{99, 11},
		true,
	},
	{
		"clocks with negative hour",
		hm{22, 40},
		hm{-2, 40},
		true,
	},
	{
		"clocks with negative hour that wraps",
		hm{17, 3},
		hm{-31, 3},
		true,
	},
	{
		"clocks with negative hour that wraps multiple times",
		hm{13, 49},
		hm{-83, 49},
		true,
	},
	{
		"clocks with minute overflow",
		hm{0, 1},
		hm{0, 1441},
		true,
	},
	{
		"clocks with minute overflow by several days",
		hm{2, 2},
		hm{2, 4322},
		true,
	},
	{
		"clocks with negative minute",
		hm{2, 40},
		hm{3, -20},
		true,
	},
	{
		"clocks with negative minute that wraps",
		hm{4, 10},
		hm{5, -1490},
		true,
	},
	{
		"clocks with negative minute that wraps multiple times",
		hm{6, 15},
		hm{6, -4305},
		true,
	},
	{
		"clocks with negative hours and minutes",
		hm{7, 32},
		hm{-12, -268},
		true,
	},
	{
		"clocks with negative hours and minutes that wrap",
		hm{18, 7},
		hm{-54, -11513},
		true,
	},
	{
		"full clock and zeroed clock",
		hm{24, 0},
		hm{0, 0},
		true,
	},
}
