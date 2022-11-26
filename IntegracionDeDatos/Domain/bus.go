package Domain

type BusLine struct {
	Line string
}

type BusStop struct {
	ID        int
	Street1   string
	Street2   string
	Street1ID string
	Street2ID string
	Location  Location
	BusLines  string
}

type BusStopDAO struct {
	ID        int
	Street1   string
	Street2   string
	Street1ID string
	Street2ID string
	Location  string
	BusLines  string
}

type Location struct {
	Type        string
	Coordinates []float64
}

type BusStopID struct {
	ID int
}
