package usecase

import (
	"fmt"

	"github.com/liampulles/go-mastermind/pkg/domain"
)

type Request [4]string

type Response struct {
	GameWon    bool
	Evaluation *domain.Evaluation
}

type HumanEngine interface {
	NewGame() error
	Evaluate(request *Request) (*Response, error)
}

type HumanEngineImpl struct {
	factory Factory
	store   GameStore
	engine  domain.Engine
}

var _ HumanEngine = &HumanEngineImpl{}

func NewHumanEngineImpl(
	factory Factory,
	store GameStore,
	engine domain.Engine,
) *HumanEngineImpl {
	return &HumanEngineImpl{
		factory: factory,
		store:   store,
		engine:  engine,
	}
}

func (he *HumanEngineImpl) NewGame() error {
	secret := he.engine.CreateSecret()
	if _, err := he.store.New(secret); err != nil {
		return fmt.Errorf("store error: %w", err)
	}
	return nil
}

func (he *HumanEngineImpl) Evaluate(request *Request) (*Response, error) {
	id, err := he.store.GetCurrentGameIdentifier()
	if err != nil {
		return nil, fmt.Errorf("could not get current game: %w", err)
	}

	won, err := he.store.IsWon(id)
	if err != nil {
		return nil, fmt.Errorf("could not get game won state: %w", err)
	}
	if won {
		return nil, fmt.Errorf("game is already won")
	}

	secret, err := he.store.GetSecret(id)
	if err != nil {
		return nil, fmt.Errorf("could not get game secret: %w", err)
	}

	guess, err := he.factory.CreateCombination(request)
	if err != nil {
		return nil, fmt.Errorf("could not understand request: %w", err)
	}

	eval := he.engine.Evaluate(secret, guess)
	if eval.IsPerfect() {
		if err := he.store.EndGame(id); err != nil {
			return nil, fmt.Errorf("could not end won game: %w", err)
		}
		return he.factory.CreateResponse(eval, true), nil
	}
	return he.factory.CreateResponse(eval, false), nil
}
