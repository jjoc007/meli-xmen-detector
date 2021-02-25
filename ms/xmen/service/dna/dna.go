package servicedna

import (
	"fmt"
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
	modeldna "github.com/jjoc007/meli-xmen-detector/xmen/model/dna"
	repositorydna "github.com/jjoc007/meli-xmen-detector/xmen/repository/dna"
	servicestats "github.com/jjoc007/meli-xmen-detector/xmen/service/stats"
	"github.com/jjoc007/meli-xmen-detector/xmen/utils"
	"github.com/pkg/errors"
)

// DNAService describes the structure a DNA service.
type DNAService interface {
	Process(*modeldna.DNA) error
	Get(string) (*modeldna.DNA, error)
}

// New creates and returns a new dna service instance
func New(rep repositorydna.DNARepository,
	statsService servicestats.StatsService) DNAService {
	return &dnaService{
		repository:             rep,
		statsService: statsService,
	}
}

type dnaService struct {
	repository             repositorydna.DNARepository
	statsService servicestats.StatsService
}

func (s *dnaService) Process(resource *modeldna.DNA) error {
	log.Logger.Debug().Msg("Validating dna")
	resource.ID = s.generateID(&resource.DNA)
	dnaStored, err := s.repository.GetByID(resource.ID)
	if err != nil {
		return errors.Wrapf(err, "Error getting dna from repository [%s]", resource.ID)
	}

	if dnaStored != nil {
		resource.IsMutant = dnaStored.IsMutant
		return nil
	}

	resource.IsMutant = s.isMutant(&resource.DNA)
	err = s.statsService.ProcessStats(resource.IsMutant)
	if err != nil {
		return errors.Wrapf(err, "Error updating stats on services [%s]", resource.ID)
	}

	err = s.repository.Create(resource)
	if err != nil {
		return errors.Wrapf(err, "Error creating dna on services [%s]", resource.ID)
	}

	return nil
}

func (s *dnaService) Get(id string) (*modeldna.DNA, error) {
	log.Logger.Debug().Msg("Getting a dna on services")
	dna, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "Error getting dna on services [%s]", id)
	}
	return dna, nil
}

func (s *dnaService) isMutant(dnaSequence *[]string) bool {
	r := len(*dnaSequence)
	c := len((*dnaSequence)[0])
	lines := uint8(0)
	validate := s.countSequences(&lines)
	for i := range *dnaSequence {
		for j := range (*dnaSequence)[i] {
			if validate(i,j, j+3 < c, dnaSequence, utils.Horizontal) ||
				validate(i,j, i+3 < r, dnaSequence, utils.Vertical)  ||
				validate(i,j, i+3 < r && j+3 < c, dnaSequence, utils.Diagonal) ||
				validate(i,j, i+3 < r && j-3 >= 0, dnaSequence, utils.InverseDiagonal) {
				return true
			}
		}
	}
	return false
}

func (s *dnaService) countSequences(lines *uint8) func(int, int, bool, *[]string, utils.Orientation) bool {
	return func(i,j int, validLimits bool, dnaSequence *[]string, orientation utils.Orientation) bool {
		if validLimits {
			*lines += s.validate(i, j ,dnaSequence, orientation)
		}

		if *lines >= 2 {
			return true
		}
		return false
	}
}

func (s *dnaService) validate(i, j int, data *[]string, orientation utils.Orientation) uint8 {
	sequence := s.getSequence(i,j,data,orientation)
	if sequence == utils.ACompleteSequence ||
		sequence == utils.TCompleteSequence ||
		sequence == utils.CCompleteSequence ||
		sequence == utils.GCompleteSequence {
		log.Logger.Debug().Msgf("Found sequence: %d, %d orietation: %v",i,j,orientation)
		return 1
	}
	return 0
}

func (s *dnaService) getSequence(i, j int, data *[]string, orientation utils.Orientation) string {
	switch orientation {
	case utils.Horizontal:
		return s.buildSequence((*data)[i][j], (*data)[i][j+1], (*data)[i][j+2], (*data)[i][j+3])
	case utils.Vertical:
		return s.buildSequence((*data)[i][j], (*data)[i+1][j], (*data)[i+2][j], (*data)[i+3][j])
	case utils.Diagonal:
		return s.buildSequence((*data)[i][j], (*data)[i+1][j+1], (*data)[i+2][j+2], (*data)[i+3][j+3])
	case utils.InverseDiagonal:
		return s.buildSequence((*data)[i][j], (*data)[i+1][j-1], (*data)[i+2][j-2], (*data)[i+3][j-3])
	}
	return ""
}

func (s *dnaService) buildSequence(letters ... uint8) string {
	sequence := ""
	for _, s := range letters{
		sequence = fmt.Sprintf("%s%s",sequence, string(s))
	}
	return sequence
}

func (s *dnaService) generateID(data *[]string) string {
	fullSequence := ""
	for _, s := range *data{
		fullSequence = fmt.Sprintf("%s%s", fullSequence, s)
	}
	return utils.Base64Encode(fullSequence)
}