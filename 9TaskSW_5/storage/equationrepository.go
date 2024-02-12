package storage

import (
	"github.com/zelinskayas/GoBasic/9TaskSW_5/internal/app/models"
	"log"
)

// instance of article repository (model inteface)
type EquationRepository struct {
	storage *Storage
}

// добавим коэффициенты квадратного уравнения
func (ar *EquationRepository) Grab(a *models.Equation) *models.Equation {
	log.Println("ar.storage.db", ar.storage.db)
	ar.storage.db.A = a.A
	ar.storage.db.B = a.B
	ar.storage.db.C = a.C
	return a
}

// Чтобы найти дискриминант, можно воспользоваться формулой и найти кол-во корней
func (ar *EquationRepository) Solve() *models.Equation {
	var nroots int
	D := ar.storage.db.B*ar.storage.db.B - 4*ar.storage.db.A*ar.storage.db.C
	if D > 0 {
		nroots = 2
	}
	if D == 0 {
		nroots = 1
	}
	if D < 0 {
		nroots = 0
	}
	ar.storage.db.Nroots = nroots
	return ar.storage.db
}
