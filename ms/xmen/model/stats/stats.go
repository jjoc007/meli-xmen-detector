package stats

type Stats struct {
	ID string `json:"id"`
	CountMutantDNA int `json:"count_mutant_dna"`
	CountHumanDNA int `json:"count_human_dna"`
	Ratio float64 `json:"ratio"`
}