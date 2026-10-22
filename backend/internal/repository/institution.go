package repository

import (
 "context"
 "errors"
 "sea/backend/internal/domain"
)

var ErrNotFound=errors.New("institution not found")
type InstitutionRepository interface {
 List(context.Context,string)([]domain.Institution,error)
 GetByID(context.Context,string)(domain.Institution,error)
}
