package Domain

type Beach struct {
	ID               int
	Name             string
	Address          string
	Coordinates      []float64
	FlagColor        string
	ChemicalBathroom bool
	HasMarket        bool
}

type BeachDAO struct {
	ID               int
	Name             string
	Address          string
	Coordinates      string
	FlagColor        string
	ChemicalBathroom bool
	HasMarket        bool
}
