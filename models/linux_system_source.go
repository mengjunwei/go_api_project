package models

const (
	SetTypeMem = "mem"
	SetTypeCpu = "cpu"
)

type SystemSetDTO struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}
