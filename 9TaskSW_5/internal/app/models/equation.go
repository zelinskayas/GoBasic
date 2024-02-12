package models

// equation model defenition
type Equation struct {
	A      int `json:"A"`
	B      int `json:"B"`
	C      int `json:"C"`
	Nroots int `json:"Nroots"`
}

func NewEquation() *Equation {
	return &Equation{}
}
