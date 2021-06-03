package usecase

import "github.com/liampulles/go-mastermind/pkg/domain"

type GameIdentifier string

type GameStore interface {
	New(secret *domain.Combination) (GameIdentifier, error)
	GetCurrentGameIdentifier() (GameIdentifier, error)
	IsWon(GameIdentifier) (bool, error)
	GetSecret(GameIdentifier) (*domain.Combination, error)
	EndGame(GameIdentifier) error
}

// Implementation is in the adapter layer
