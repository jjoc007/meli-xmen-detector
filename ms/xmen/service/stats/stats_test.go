package servicestats

import (
	statmodel "github.com/jjoc007/meli-xmen-detector/xmen/model/stats"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Test New", func(t *testing.T) {
		got := New(&statsRepositoryMock{})
		if got == nil {
			t.Errorf("NewRepository() = %v is nil", got)
		}
	})
}

func Test_statsService_ProcessStats(t *testing.T) {
	isMutant := true
	s := New(&statsRepositoryMock{})
	t.Run("ProcessStats is mutatnt", func(t *testing.T) {
		err := s.ProcessStats(isMutant)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}
	})
	isMutant = false
	t.Run("ProcessStats is not mutant", func(t *testing.T) {
		err := s.ProcessStats(isMutant)
		if err != nil {
			t.Errorf("Update() error = %v", err)
		}
	})
}

func Test_statsService_Get(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		s := New(&statsRepositoryMock{})
		res, err := s.Get()
		if err != nil {
			t.Errorf("Get() error = %v", err)
		}
		if res.ID != mockStats.ID {
			t.Errorf("GetByID() id = %v, want %v", res.ID, mockStats.ID)
		}
		if res.CountHumanDNA != mockStats.CountHumanDNA {
			t.Errorf("GetByID() CountHumanDNA = %v, want %v", res.CountHumanDNA, mockStats.CountHumanDNA)
		}
		if res.CountMutantDNA != mockStats.CountMutantDNA {
			t.Errorf("GetByID() CountMutantDNA = %v, want %v", res.CountMutantDNA, mockStats.CountMutantDNA)
		}
		if res.Ratio != mockStats.Ratio {
			t.Errorf("GetByID() Ratio = %v, want %v", res.Ratio, mockStats.Ratio)
		}
	})
}

var mockStats = &statmodel.Stats{
	ID:             "1",
	CountMutantDNA: 10,
	CountHumanDNA:  10,
	Ratio:          10,
}

type statsRepositoryMock struct {}
func (m *statsRepositoryMock) Update(resource *statmodel.Stats) error {
	return nil
}

func (m *statsRepositoryMock) GetByID(id string) (*statmodel.Stats, error) {
	return 	mockStats , nil
}