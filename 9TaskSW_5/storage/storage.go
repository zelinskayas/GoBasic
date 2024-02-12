package storage

import "github.com/zelinskayas/GoBasic/9TaskSW_5/internal/app/models"

// instance of storage
type Storage struct {
	config *Config
	//subfield for repo interfacing (model user)
	equationRepository *EquationRepository
	db                 *models.Equation
}

func New(config *Config, db *models.Equation) *Storage {
	return &Storage{
		config: config,
		db:     db,
	}
}

// public repo for equation
func (s *Storage) Equation() *EquationRepository {
	if s.equationRepository != nil {
		return s.equationRepository
	}
	s.equationRepository = &EquationRepository{
		storage: s,
	}
	return s.equationRepository
}
