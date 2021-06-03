package usecase

import "github.com/liampulles/go-mastermind/pkg/domain"

type GameStore interface {
	New(secret *domain.Combination) error
	IsWon() (bool, error)
	GetSecret() (*domain.Combination, error)
	EndGame() error
}

// Implementation is in the adapter layer
