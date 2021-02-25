package servicedna

import (
	dnamodel "github.com/jjoc007/meli-xmen-detector/xmen/model/dna"
	statmodel "github.com/jjoc007/meli-xmen-detector/xmen/model/stats"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Test New", func(t *testing.T) {
		got := New(&dnaRepositoryMock{}, &statsServiceMock{})
		if got == nil {
			t.Errorf("NewRepository() = %v is nil", got)
		}
	})
}

func Test_statsService_Process(t *testing.T) {
	s := New(&dnaRepositoryMock{}, &statsServiceMock{})
	mockDnaCases := &dnamodel.DNA{
		ID:       "1",
		DNA: []string{
			"CCCCGA",
			"CAGTCC",
			"GTTAGT",
			"CGAAGG",
			"TGTCTA",
			"TCCCCC",
		},
		IsMutant: false,
	}
	t.Run("Process is mutant with horizontal orientation", func(t *testing.T) {
		err := s.Process(mockDnaCases)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}

		if !mockDnaCases.IsMutant{
			t.Errorf("Process() IsMutant = %v, want %v", mockDnaCases.IsMutant, true)
		}
	})

	mockDnaCases.DNA = []string{
		"CCCCGA",
		"CAGTCC",
		"GATAGT",
		"CAAAGG",
		"TATCTA",
		"TCCTGC",
	}

	t.Run("Process is mutant with vertical orientation", func(t *testing.T) {
		err := s.Process(mockDnaCases)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}

		if !mockDnaCases.IsMutant{
			t.Errorf("Process() IsMutant = %v, want %v", mockDnaCases.IsMutant, true)
		}
	})

	mockDnaCases.DNA = []string{
		"CCCCGA",
		"CAGTCC",
		"GTAAGT",
		"CAAAGG",
		"TCTCAA",
		"TCCTGC",
	}

	t.Run("Process is mutant with diagonal orientation", func(t *testing.T) {
		err := s.Process(mockDnaCases)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}

		if !mockDnaCases.IsMutant{
			t.Errorf("Process() IsMutant = %v, want %v", mockDnaCases.IsMutant, true)
		}
	})

	mockDnaCases.DNA = []string{
		"CCCCGA",
		"CAGTAC",
		"GTGAGT",
		"CAATGG",
		"TCTCAA",
		"TCCTGC",
	}

	t.Run("Process is mutant with inverse diagonal orientation", func(t *testing.T) {
		err := s.Process(mockDnaCases)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}

		if !mockDnaCases.IsMutant{
			t.Errorf("Process() IsMutant = %v, want %v", mockDnaCases.IsMutant, true)
		}
	})

	mockDnaCases.DNA = []string{
		"CCACGA",
		"CAGTTC",
		"GTGAGT",
		"CAGTGG",
		"TCTCAA",
		"TCCTGC",
	}

	t.Run("Process is not mutant", func(t *testing.T) {
		err := s.Process(mockDnaCases)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}

		if mockDnaCases.IsMutant{
			t.Errorf("Process() IsMutant = %v, want %v", mockDnaCases.IsMutant, false)
		}
	})
}

var mockStats = &statmodel.Stats{
	ID:             "1",
	CountMutantDNA: 10,
	CountHumanDNA:  10,
	Ratio:          10,
}

type statsServiceMock struct {}
func (m *statsServiceMock)Get() (*statmodel.Stats, error) {
	return mockStats, nil
}

func (m *statsServiceMock) ProcessStats(isMutant bool) error {
	return nil
}

type dnaRepositoryMock struct {}
func (m *dnaRepositoryMock) Create(resource *dnamodel.DNA) (err error) {
	return nil
}

func (m *dnaRepositoryMock) GetByID(id string) (dna *dnamodel.DNA, err error) {
	return nil, nil
}