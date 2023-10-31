package measurement

type Dimensions struct {
	Width    Length `json:"width,omitempty"`
	Heigth   Length `json:"heigth,omitempty"`
	Depth    Length `json:"depth,omitempty"`
	Length   Length `json:"length,omitempty"`
	Radius   Length `json:"radius,omitempty"`
	Diameter Length `json:"diameter,omitempty"`
}
