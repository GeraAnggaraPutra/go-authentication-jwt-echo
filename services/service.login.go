package services

import (
	"go-auth-jwt/repository"
)

type ServiceLogin interface {
	CheckLogin(username, password string) (bool, error)
}

type serviceLogin struct {
	repository repository.RepositoryLogin
}

func NewServiceLogin(repository repository.RepositoryLogin) *serviceLogin {
	return &serviceLogin{repository}
}

func (s *serviceLogin) CheckLogin(username, password string) (bool, error) {
	res, err := s.repository.CheckLogin(username, password)
	return res, err
}