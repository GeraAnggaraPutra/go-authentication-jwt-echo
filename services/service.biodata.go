package services

import (
	"go-auth-jwt/entity"
	"go-auth-jwt/repository"
)

type Service interface {
	GetAll() ([]*entity.Biodata, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]*entity.Biodata, error) {
	biodatas, err := s.repository.GetAll()
	return biodatas, err
}