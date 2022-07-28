package service

import (
	"github.com/winnix/shifting/api/domain"
	"gorm.io/gorm"
)

type AccountTokenService interface {
	GetByApaleoToken(apaOneToken string) (*domain.AccountToken, error)
	GetCurrent() (*domain.AccountToken, error)
}

type accountTokenService struct {
	db          *gorm.DB
	accountCode string
}

func NewAccountTokenService(db *gorm.DB, accountCode string) AccountTokenService {
	return &accountTokenService{
		db:          db,
		accountCode: accountCode,
	}
}

func (s *accountTokenService) GetByApaleoToken(apaOneToken string) (*domain.AccountToken, error) {
	res := domain.AccountToken{}
	err := s.db.
		Where(&domain.AccountToken{AccountCode: s.accountCode, ApaleoOneToken: apaOneToken}).
		First(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *accountTokenService) GetCurrent() (*domain.AccountToken, error) {
	res := domain.AccountToken{}
	err := s.db.
		Where(&domain.AccountToken{AccountCode: s.accountCode}).
		FirstOrCreate(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}
