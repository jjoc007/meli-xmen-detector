package servicestats

import (
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
	modelstats "github.com/jjoc007/meli-xmen-detector/xmen/model/stats"
	repositorystats "github.com/jjoc007/meli-xmen-detector/xmen/repository/stats"
	"github.com/pkg/errors"
)

// StatsService describes the structure a Stats service.
type StatsService interface {
	ProcessStats(bool) error
	Get() (*modelstats.Stats, error)
}

// New creates and returns a new stats service instance
func New(rep repositorystats.StatsRepository) StatsService {
	return &statsService{
		repository:             rep,
	}
}

type statsService struct {
	repository             repositorystats.StatsRepository
}

func (s *statsService) ProcessStats(isMutant bool) error {
	log.Logger.Debug().Msg("Updating record")
	stats := &modelstats.Stats{}
	stats, err :=s.repository.GetByID("1")
	if err != nil {
		return errors.Wrapf(err, "Error getting stats from repository [%s]","1")
	}

	stats.ID = "1"
	s.updateMutant(stats, isMutant)
	err = s.repository.Update(stats)
	if err != nil {
		return errors.Wrap(err, "Error updating stats")
	}

	return nil
}

func (s *statsService) Get() (*modelstats.Stats, error) {
	log.Logger.Debug().Msg("Getting a dna on services")
	stats, err :=s.repository.GetByID("1")
	if err != nil {
		return nil, errors.Wrapf(err, "Error getting stats from repository [%s]","1")
	}
	return stats, nil
}


func (s *statsService) updateMutant(stats *modelstats.Stats, isMutant bool){
	if isMutant {
		stats.CountMutantDNA = stats.CountMutantDNA + 1
	} else {
		stats.CountHumanDNA = stats.CountHumanDNA + 1
	}

	if stats.CountHumanDNA == 0 {
		stats.Ratio =float64(stats.CountMutantDNA)
	}else {
		stats.Ratio = float64(stats.CountMutantDNA) / float64(stats.CountHumanDNA)
	}

}