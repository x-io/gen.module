package dbs

type Operator struct {
	ID   String `json:"id,omitempty"`
	Name String `json:"name,omitempty"`

	Depot String `json:"depot,omitempty"`
}
