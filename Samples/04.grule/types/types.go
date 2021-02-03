package types

type Gender string

const (
	Male    = "Male"
	Female  = "Female"
	Unknown = "Unknown"
)

type CustomerFact struct {
	Name       string
	Gender     Gender
	Age        int16
	City       string
	IsDisabled bool
}
