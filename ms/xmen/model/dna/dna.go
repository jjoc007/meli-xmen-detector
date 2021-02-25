package dna

type DNA struct {
	ID string `json:"id"`
	DNA []string `json:"dna"`
	IsMutant bool `json:"is_mutant"`
}